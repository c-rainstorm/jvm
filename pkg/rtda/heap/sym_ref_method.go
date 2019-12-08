package heap

import "jvm/pkg/classfile"

type MethodSymRef struct {
	ClassMemberSymRef
	method *Method
}

func newMethodSymRef(cp *ConstantPool, cfField *classfile.ConstantMethodRefInfo) Constant {
	fieldSymRef := &MethodSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
