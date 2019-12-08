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
	if this.isPublic() {
		// public 字段任意类都可以访问
		return true
	}

	currentClass := this.class

	if this.isProtected() {
	 	// protected 成员只能当前类、当前类的子类，同一个包下面的类可访问
		return currentClass == accessClass || accessClass.isSubClassOf(currentClass) ||
			currentClass.packageName() == accessClass.packageName()
	}

	if !this.isPrivate() {
		// default 成员同一个包下的可以访问
		return currentClass.packageName() == accessClass.packageName()
	}

	// private 成员只能由当前类访问
	return currentClass == accessClass
}

func (this *ClassMember) isPublic() bool {
	return this.hasFlag(ACC_PUBLIC)
}

func (this *ClassMember) isProtected() bool {
	return this.hasFlag(ACC_PROTECTED)
}

func (this *ClassMember) isPrivate() bool {
	return this.hasFlag(ACC_PRIVATE)
}

func (this *ClassMember) Class() *Class {
	return this.class
}

func (this *ClassMember) Descriptor() string {
	return this.descriptor
}
