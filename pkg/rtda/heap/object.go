package heap

import (
	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

type BaseObject struct {
	class *ClassObject
}

type Object interface {
	Class() *ClassObject

	IsInstanceOf(targetClass *ClassObject) bool
}

func (this *BaseObject) Class() *ClassObject {
	return this.class
}

func (this *BaseObject) IsInstanceOf(targetClass *ClassObject) bool {
	return this.isAssignableTo(targetClass)
}

func (this *BaseObject) isAssignableTo(targetClass *ClassObject) bool {
	if this.class == targetClass {
		return true
	}

	if this.class.IsArray() {
		if targetClass.IsInterface() {
			return global.JavaIOSerializable == targetClass.Name() ||
				global.JavaLangCloneable == targetClass.Name()
		} else if targetClass.IsArray() {
			thisEle := this.class.ElementClass()
			targetEle := targetClass.ElementClass()
			return thisEle == targetEle || thisEle.IsSubClassOf(targetEle)
		} else {
			return global.JavaLangObject == targetClass.Name()
		}
	} else if this.class.IsInterface() {
		if targetClass.IsInterface() {
			return this.class.IsSubInterfaceOf(targetClass)
		} else {
			return global.JavaLangObject == targetClass.Name()
		}
	} else {
		if targetClass.IsInterface() {
			return this.class.IsImplClassOf(targetClass)
		} else {
			return this.class.IsSubClassOf(targetClass)
		}
	}
}
