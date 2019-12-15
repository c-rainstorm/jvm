package heap

import "jvm/pkg/logger"

var log = logger.NewLogrusLogger()

type Object struct {
	class  *Class
	fields Slots
}

func (this *Object) FieldSlots() Slots {
	return this.fields
}

func (this *Object) IsInstanceOf(targetClass *Class) bool {
	return this.isAssignableTo(targetClass)
}

func (this *Object) isAssignableTo(targetClass *Class) bool {
	if this.class == targetClass {
		return true
	}
	if targetClass.IsInterface() {
		return this.class.IsImplClassOf(targetClass)
	} else {
		return this.class.IsSubClassOf(targetClass)
	}
}

func (this *Object) Class() *Class {
	return this.class
}
