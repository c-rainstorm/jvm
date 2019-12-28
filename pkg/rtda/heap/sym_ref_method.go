package heap

import (
	"jvm/pkg/classfile"
)

type MethodSymRef struct {
	ClassMemberSymRef
	method *Method
}

type IMethodSymRef interface {
	ResolvedMethod() *Method
}

func (this *MethodSymRef) ResolvedMethod() *Method {
	if this.method == nil {
		this.resolveMethodRef()
	}

	return this.method
}

func (this *MethodSymRef) resolveMethodRef() {
	currentClass := this.cp.class
	classHasThisMethod := this.ResolvedClass()
	method := lookupMethod(classHasThisMethod, this.name, this.descriptor)

	if method == nil {
		panic("java.lang.NoSuchMethodError")
	}

	if !method.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	this.method = method
}

func newMethodSymRef(cp *ConstantPool, cfField *classfile.ConstantMethodRefInfo) Constant {
	fieldSymRef := &MethodSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
