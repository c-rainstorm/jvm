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

func (this *ClassMember) isAccessibleTo(accessClass *Class) bool {
	if this.IsPublic() {
		// public 字段任意类都可以访问
		return true
	}

	currentClass := this.class

	if this.IsProtected() {
		// protected 成员只能当前类、当前类的子类，同一个包下面的类可访问
		return currentClass == accessClass || accessClass.IsSubClassOf(currentClass) ||
			currentClass.PackageName() == accessClass.PackageName()
	}

	if !this.IsPrivate() {
		// default 成员同一个包下的可以访问
		return currentClass.PackageName() == accessClass.PackageName()
	}

	// private 成员只能由当前类访问
	return currentClass == accessClass
}

func (this *ClassMember) IsPublic() bool {
	return this.hasFlag(ACC_PUBLIC)
}

func (this *ClassMember) IsProtected() bool {
	return this.hasFlag(ACC_PROTECTED)
}

func (this *ClassMember) IsPrivate() bool {
	return this.hasFlag(ACC_PRIVATE)
}

func (this *Method) IsStatic() bool {
	return this.hasFlag(ACC_STATIC)
}

func (this *ClassMember) Class() *Class {
	return this.class
}

func (this *ClassMember) Descriptor() string {
	return this.descriptor
}

func (this *ClassMember) Name() string {
	return this.name
}
