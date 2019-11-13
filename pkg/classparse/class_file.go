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

func (this *ClassFile) read(reader *ClassReader) *ClassFile {
	this.readMagic(reader)
	return this
}

func (this *ClassFile) readMagic(reader *ClassReader) {
	this.magic = reader.readUint32()
	if global.Verbose {
		log.Infof("magic parsed: %v", this.Magic())
	}
	if this.magic != global.JavaClassFileMagic {
		log.Panicf("java.lang.ClassFormatError: magic! expected: %s, actual: %s",
			global.JavaClassFileMagic, this.magic)
	}
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

func (this *ClassFile) String() string {
	return fmt.Sprintf("ClassFile{magic: %s}", this.Magic())
}

func (this *ClassFile) Magic() string {
	return fmt.Sprintf("0x%X", this.magic)
}
