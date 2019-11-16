package classparse

import (
	"fmt"

	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

type ConstantPool struct {
}

type MemberInfo struct {
}

type AttributeInfo struct {
}

type ClassFile struct {
	// 魔术，class 文件固定 0xCAFEBABE
	magic        uint32
	minorVersion uint16
	majorVersion uint16
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	supperClass  uint16
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
	return fmt.Sprintf("ClassFile{magic: %s, version: %s}",
		this.Magic(), this.Version())
}

func (this *ClassFile) Magic() string {
	return fmt.Sprintf("0x%X", this.magic)
}

func (this *ClassFile) Version() interface{} {
	return fmt.Sprintf("%v.%v", this.majorVersion, this.minorVersion)
}
