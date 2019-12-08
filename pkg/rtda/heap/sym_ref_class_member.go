package heap

import "jvm/pkg/classfile"

type ClassMemberSymRef struct {
	SymbolicRef
	name       string
	descriptor string
}

func (this *ClassMemberSymRef) Name() string {
	return this.name
}

func (this *ClassMemberSymRef) Descriptor() string {
	return this.descriptor
}

func (this *ClassMemberSymRef) copy(cfMember *classfile.ConstantMemberRefInfo) {
	this.className = cfMember.ClassName()
	this.name, this.descriptor = cfMember.NameAndDescriptor()
}
