package heap

import "jvm/pkg/classfile"

type FieldSymRef struct {
	ClassMemberSymRef
	field *Field
}

func (this *FieldSymRef) ResolvedField() *Field {
	if this.field == nil {
		this.resolveFieldRef()
	}

	return this.field
}

func (this *FieldSymRef) resolveFieldRef() {
	currentClass := this.cp.class
	classHasThisField := this.ResolvedClass()
	field := classHasThisField.lookupField(this.name, this.descriptor)

	if field == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !field.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	this.field = field
}

func newFieldSymRef(cp *ConstantPool, cfField *classfile.ConstantFieldRefInfo) Constant {
	fieldSymRef := &FieldSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
