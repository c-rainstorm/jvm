package misc

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func init() {
	native.MethodRegistry.Registry("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	// 添加一个元素到 savedProps
	VM := frame.Method().Class()
	savedProps := VM.GetField(nil, "savedProps", "Ljava/util/Properties;").(*heap.NormalObject)

	operandStack := frame.OperandStack()
	operandStack.PushRef(savedProps)
	operandStack.PushRef(heap.JString(VM.ClassLoader(), "key"))
	operandStack.PushRef(heap.JString(VM.ClassLoader(), "value"))

	// Object setProperty(String key, String value) -> (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;
	setPropertyMethod := savedProps.Class().GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	base.InvokeMethod(frame, setPropertyMethod)
}
