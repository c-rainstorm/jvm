package invoke

import (
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func SetProperty(thread *rtda.Thread, properties *heap.NormalObject, key, value string) {
	thread.PushFrame(thread.NewFrame(heap.VoidVirtualMethod(3, 0)))

	frame := thread.TopFrame()
	propertiesClass := properties.Class()
	operandStack := frame.OperandStack()
	operandStack.PushRef(properties)
	operandStack.PushRef(heap.JString(propertiesClass.ClassLoader(), key))
	operandStack.PushRef(heap.JString(propertiesClass.ClassLoader(), value))

	// Object setProperty(String key, String value) -> (Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;
	setPropertyMethod := propertiesClass.GetMethod("setProperty", "(Ljava/lang/String;Ljava/lang/String;)Ljava/lang/Object;")

	InvokeMethod(frame, setPropertyMethod)
}
