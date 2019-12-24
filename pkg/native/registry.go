package native

import (
	"strings"

	"jvm/pkg/rtda"
)

// Method
type Method func(frame *rtda.Frame)

type TMethodRegistry map[string]Method

var MethodRegistry = TMethodRegistry{}

func (this TMethodRegistry) Registry(classname, methodName, methodDescriptor string, method Method) {
	this[this.getKey(classname, methodName, methodDescriptor)] = method
}

func (this TMethodRegistry) Find(classname, methodName, methodDescriptor string) Method {
	if method, ok := this[this.getKey(classname, methodName, methodDescriptor)]; ok {
		return method
	}

	if methodDescriptor == "()V" && methodName == "registerNatives" {
		return doNothing
	}

	return nil
}

func doNothing(frame *rtda.Frame) {
}

func (this TMethodRegistry) getKey(classname string, name string, descriptor string) string {
	return strings.Join([]string{classname, name, descriptor}, "~")
}
