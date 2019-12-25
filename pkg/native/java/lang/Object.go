package lang

import (
	"jvm/pkg/global"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry(global.JavaLangObject, "getClass", "()Ljava/lang/Class;", getClass)
	native.MethodRegistry.Registry(global.JavaLangObject, "hashCode", "()I", hashCode)
}

func hashCode(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	frame.OperandStack().PushInt(this.HashCode())
}

func getClass(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	this := localVars.GetThis()
	frame.OperandStack().PushRef(this.Class())
}
