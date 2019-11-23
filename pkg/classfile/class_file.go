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
	fields       []*FieldMemberInfo
	methods      []*MethodMemberInfo
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
	this.fields = make([]*FieldMemberInfo, memberCount)
	for i := range this.fields {
		this.fields[i] = &FieldMemberInfo{MemberInfo{}}
		this.fields[i].read(reader, this.constantPool)
	}
}

func (this *ClassFile) readMethods(reader *ClassReader) {
	memberCount := reader.readUnit16()
	this.methods = make([]*MethodMemberInfo, memberCount)
	for i := range this.methods {
		this.methods[i] = &MethodMemberInfo{MemberInfo{}}
		this.methods[i].read(reader, this.constantPool)
	}
}

func (this *ClassFile) String() string {
	var classFileInfo = make(map[string]string)
	classFileInfo["magic"] = this.Magic()
	classFileInfo["version"] = this.Version()
	classFileInfo["accessFlags"] = this.AccessFlag()
	classFileInfo["thisClass"] = this.ThisClass()
	classFileInfo["superClass"] = this.SuperClass()
	classFileInfo["constant pool"] = this.constantPool.String()
	classFileInfo["interfaces"] = this.Interfaces()
	classFileInfo["fields"] = this.Fields()
	classFileInfo["methods"] = this.Methods()
	return fmt.Sprintf("ClassFile%v", classFileInfo)
}

func (this *ClassFile) Magic() string {
	return fmt.Sprintf("0x%X", this.magic)
}

func (this *ClassFile) Version() string {
	return fmt.Sprintf("%v.%v", this.majorVersion, this.minorVersion)
}

func (this *ClassFile) AccessFlag() string {
	builder := strings.Builder{}
	builder.WriteString(this.checkAccessFlag(accClassPublic, global.KeywordPublic))
	builder.WriteString(this.checkAccessFlag(accClassFinal, global.KeywordFinal))
	builder.WriteString(this.checkAccessFlag(accClassAbstract, global.KeywordAbstract))
	if accClassInterface&this.accessFlags != 0 {
		if accClassAnnotation&this.accessFlags != 0 {
			builder.WriteString(global.KeywordAnnotation)
		} else {
			builder.WriteString(global.KeywordInterface)
		}
		builder.WriteString(global.Space)
	}
	builder.WriteString(this.checkAccessFlag(accClassEnum, global.KeywordEnum))
	builder.WriteString(this.checkAccessFlag(accClassSynthetic, global.AccGenerated))

	return builder.String()
}

func (this *ClassFile) checkAccessFlag(targetFlag uint16, targetKeyword string) string {
	if targetFlag&this.accessFlags != 0 {
		return fmt.Sprintf("%v ", targetKeyword)
	}
	return global.EmptyString
}

func (this *ClassFile) ThisClass() string {
	return fmt.Sprintf("%v", this.constantPool[this.thisClass])
}

func (this *ClassFile) SuperClass() string {
	return fmt.Sprintf("%v", this.constantPool[this.superClass])
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

func (this *ClassFile) Fields() string {
	builder := strings.Builder{}
	builder.WriteString("Fields{")

	for i := range this.fields {
		builder.WriteString(strings.Join([]string{fmt.Sprintf("%v", this.fields[i]), ",\n"}, ""))
	}

	builder.WriteString("}")

	return builder.String()
}

func (this *ClassFile) Methods() string {
	builder := strings.Builder{}
	builder.WriteString("Methods{")

	for i := range this.methods {
		builder.WriteString(strings.Join([]string{fmt.Sprintf("%v", this.methods[i]), ",\n"}, ""))
	}

	builder.WriteString("}")

	return builder.String()
}