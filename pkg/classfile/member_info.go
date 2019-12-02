package classfile

import (
	"fmt"

	"jvm/pkg/global"
)

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func (this *MemberInfo) read(reader *ClassReader, cp ConstantPool) {
	this.cp = cp
	this.accessFlags = reader.readUnit16()
	this.nameIndex = reader.readUnit16()
	this.descriptorIndex = reader.readUnit16()
	this.attributes = readAttributes(reader, cp)
}

func (this *MemberInfo) checkAccessFlag(targetFlag uint16, targetKeyword string) string {
	if targetFlag&this.accessFlags != 0 {
		return fmt.Sprintf("%v ", targetKeyword)
	}
	return global.EmptyString
}

func (this *MemberInfo) Name() string {
	return this.cp[this.nameIndex].(*ConstantUtf8Info).val
}

func (this *MemberInfo) Descriptor() string {
	return this.cp[this.descriptorIndex].(*ConstantUtf8Info).val
}
