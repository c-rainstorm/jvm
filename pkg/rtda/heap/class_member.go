package heap

import "jvm/pkg/classfile"

type ClassMember struct {
	accessFlags uint16
	name        string
	descriptor  string
	class       *Class
}

func (this *ClassMember) copy(cfMemberInfo *classfile.MemberInfo) {
	this.accessFlags = cfMemberInfo.AccessFlags()
	this.name = cfMemberInfo.Name()
	this.descriptor = cfMemberInfo.Descriptor()
}

func (this *ClassMember) hasFlag(flag uint16) bool {
	return (this.accessFlags & flag) != 0
}
