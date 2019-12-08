package classfile

import (
	"fmt"
	"strings"

	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

const (
	accClassPublic     uint16 = 0x0001
	accClassFinal      uint16 = 0x0010
	accClassSuper      uint16 = 0x0020
	accClassInterface  uint16 = 0x0200
	accClassAbstract   uint16 = 0x0400
	accClassSynthetic  uint16 = 0x1000
	accClassAnnotation uint16 = 0x2000
	accClassEnum       uint16 = 0x4000
)

type ClassFile struct {
	// 魔数，class 文件固定 0xCAFEBABE
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classBytes []byte) *ClassFile {
	classReader := &ClassReader{
		data:  classBytes,
		index: 0,
	}

	return newClassFile(classReader)
}

func newClassFile(reader *ClassReader) *ClassFile {
	classFile := &ClassFile{}

	return classFile.read(reader)
}

func (this *ClassFile) read(reader *ClassReader) *ClassFile {
	this.readMagic(reader)
	this.readVersion(reader)
	this.readConstantPool(reader)
	this.accessFlags = reader.readUnit16()
	this.thisClass = reader.readUnit16()
	this.superClass = reader.readUnit16()
	this.interfaces = reader.readUint16s(reader.readUnit16())
	this.readFields(reader)
	this.readMethods(reader)
	this.attributes = readAttributes(reader, this.constantPool)
	return this
}

func (this *ClassFile) readMagic(reader *ClassReader) {
	this.magic = reader.readUint32()
	if global.Verbose {
		log.Infof("parsed magic: %v", this.Magic())
	}
	if this.magic != global.JavaClassFileMagic {
		log.Panicf("java.lang.ClassFormatError: magic! expected: %s, actual: %s",
			global.JavaClassFileMagic, this.magic)
	}
}

func (this *ClassFile) readVersion(reader *ClassReader) {
	this.minorVersion = reader.readUnit16()
	this.majorVersion = reader.readUnit16()
	if global.Verbose {
		log.Infof("parsed version: %v", this.Version())
	}

	switch this.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if this.minorVersion == 0 {
			return
		}
	}

	log.Panic("java.lang.UnsupportedClassVersionError!")
}

func (this *ClassFile) readConstantPool(reader *ClassReader) {
	this.constantPool = readConstantPool(reader)
}

func (this *ClassFile) readFields(reader *ClassReader) {
	memberCount := reader.readUnit16()
	this.fields = make([]*MemberInfo, memberCount)
	for i := range this.fields {
		this.fields[i] = &MemberInfo{}
		this.fields[i].read(reader, this.constantPool)
	}
}

func (this *ClassFile) readMethods(reader *ClassReader) {
	memberCount := reader.readUnit16()
	this.methods = make([]*MemberInfo, memberCount)
	for i := range this.methods {
		this.methods[i] = &MemberInfo{}
		this.methods[i].read(reader, this.constantPool)
	}
}

func (this *ClassFile) Magic() string {
	return fmt.Sprintf("0x%X", this.magic)
}

func (this *ClassFile) Version() string {
	return fmt.Sprintf("%v.%v", this.majorVersion, this.minorVersion)
}

func (this *ClassFile) AccessFlag() uint16 {
	return this.accessFlags
}

func (this *ClassFile) checkAccessFlag(targetFlag uint16, targetKeyword string) string {
	if targetFlag&this.accessFlags != 0 {
		return fmt.Sprintf("%v ", targetKeyword)
	}
	return global.EmptyString
}

func (this *ClassFile) ClassName() string {
	return this.getClassName(this.thisClass)
}

func (this *ClassFile) SuperClassName() string {
	if this.superClass > 0 {
		return this.getClassName(this.superClass)
	}
	return global.EmptyString
}

func (this *ClassFile) Interfaces() string {
	builder := strings.Builder{}
	builder.WriteString("{")
	length := len(this.interfaces)
	for i := 0; i < length; i++ {
		constantInfo := this.constantPool[this.interfaces[i]]

		builder.WriteString(strings.Join([]string{fmt.Sprintf("%v", constantInfo), ","}, ""))
	}
	builder.WriteString("}")

	return builder.String()
}

func (this *ClassFile) Fields() []*MemberInfo {
	return this.fields
}

func (this *ClassFile) Methods() []*MemberInfo {
	return this.methods
}

func (this *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(this.interfaces))
	for i, cpIndex := range this.interfaces {
		interfaceNames[i] = this.constantPool[cpIndex].(*ConstantUtf8Info).val
	}
	return interfaceNames
}

func (this *ClassFile) ConstantPool() ConstantPool {
	return this.constantPool
}

func (this *ClassFile) getClassName(classIndex uint16) string {
	classNameIndex := this.constantPool[classIndex].(*ConstantClassInfo).index
	return this.constantPool[classNameIndex].(*ConstantUtf8Info).val
}
