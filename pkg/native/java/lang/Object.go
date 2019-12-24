package lang

import (
	"jvm/pkg/global"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry(global.JavaLangObject, "getClass", "()Ljava/lang/Class;", getClass)
}

func getClass(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	this := localVars.GetThis()
	frame.OperandStack().PushRef(this.Class())
}
