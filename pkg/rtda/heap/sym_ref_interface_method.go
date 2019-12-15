package heap

import "jvm/pkg/classfile"

type InterfaceMethodSymRef struct {
	MethodSymRef
}

func (this *InterfaceMethodSymRef) ResolvedInterfaceMethod() *Method {
	if this.method == nil {
		this.resolveInterfaceMethodRef()
	}

	return this.method
}

func (this *InterfaceMethodSymRef) resolveInterfaceMethodRef() {
	currentClass := this.cp.class
	interfaceHasThisMethod := this.ResolvedClass()

	if !interfaceHasThisMethod.IsInterface() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	method := lookupMethodInInterface(interfaceHasThisMethod, this.name, this.descriptor)

	if method == nil {
		panic("java.lang.NoSuchFieldError")
	}

	if !method.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	this.method = method
}

func newInterfaceMethodSymRef(cp *ConstantPool, cfField *classfile.ConstantInterfaceMethodRefInfo) Constant {
	fieldSymRef := &InterfaceMethodSymRef{}
	fieldSymRef.cp = cp
	fieldSymRef.copy(&cfField.ConstantMemberRefInfo)
	return fieldSymRef
}
