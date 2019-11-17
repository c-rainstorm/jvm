package classfile

import (
	"fmt"
	"strings"

	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

const (
	acc_public     uint16 = 0x0001
	acc_final      uint16 = 0x0010
	acc_super      uint16 = 0x0020
	acc_interface  uint16 = 0x0200
	acc_abstract   uint16 = 0x0400
	acc_synthetic  uint16 = 0x1000
	acc_annotation uint16 = 0x2000
	acc_enum       uint16 = 0x4000
)

type MemberInfo struct {
}

type AttributeInfo struct {
}

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
	this.constantPool = readConstantPool(reader)
	this.accessFlags = reader.readUnit16()
	this.thisClass = reader.readUnit16()
	this.superClass = reader.readUnit16()
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

func (this *ClassFile) String() string {
	var classFileInfo = make(map[string]string)
	classFileInfo["magic"] = this.Magic()
	classFileInfo["version"] = this.Version()
	classFileInfo["accessFlags"] = this.AccessFlag()
	classFileInfo["thisClass"] = this.ThisClass()
	classFileInfo["superClass"] = this.SuperClass()
	classFileInfo["constant pool"] = this.constantPool.String()
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
	builder.WriteString(this.checkAccessFlag(acc_public, global.KeywordPublic))
	builder.WriteString(this.checkAccessFlag(acc_final, global.KeywordFinal))
	builder.WriteString(this.checkAccessFlag(acc_abstract, global.KeywordAbstract))
	if acc_interface&this.accessFlags != 0 {
		if acc_annotation&this.accessFlags != 0 {
			builder.WriteString(global.KeywordAnnotation)
		} else {
			builder.WriteString(global.KeywordInterface)
		}
		builder.WriteString(global.Space)
	}
	builder.WriteString(this.checkAccessFlag(acc_enum, global.KeywordEnum))
	builder.WriteString(this.checkAccessFlag(acc_synthetic, global.AccGenerated))

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
