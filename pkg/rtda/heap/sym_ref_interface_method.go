package heap

import "jvm/pkg/classfile"

type InterfaceMethodSymRef struct {
	MethodSymRef
}

func newInterfaceMethodSymRef(cp *ConstantPool, cfField *classfile.ConstantInterfaceMethodRefInfo) Constant {
	fieldSymRef := &InterfaceMethodSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
