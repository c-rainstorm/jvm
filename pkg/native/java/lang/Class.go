package lang

import (
	"jvm/pkg/global"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func init() {
	native.MethodRegistry.Registry(global.JavaLangClass, "getPrimitiveClass", "(Ljava/lang/String;)Ljava/lang/Class;", getPrimitiveClass)
	native.MethodRegistry.Registry(global.JavaLangClass, "getName0", "()Ljava/lang/String;", getName0)
	native.MethodRegistry.Registry(global.JavaLangClass, "desiredAssertionStatus0", "(Ljava/lang/Class;)Z", desiredAssertionStatus0)
}

func desiredAssertionStatus0(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func getName0(frame *rtda.Frame) {
	this := frame.LocalVars().GetThis()
	class := this.(*heap.ClassObject)
	frame.OperandStack().PushRef(heap.JString(class.ClassLoader(), class.JavaName()))
}

func getPrimitiveClass(frame *rtda.Frame) {
	loader := frame.Method().Class().ClassLoader()
	name := frame.LocalVars().GetNormalObject(0)
	GoName := heap.GoString(name)
	frame.OperandStack().PushRef(loader.LoadClass(GoName))
}
