package heap

import "jvm/pkg/classfile"

type FieldSymRef struct {
	ClassMemberSymRef
	field *Field
}

func newFieldSymRef(cp *ConstantPool, cfField *classfile.ConstantFieldRefInfo) Constant {
	fieldSymRef := &FieldSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
