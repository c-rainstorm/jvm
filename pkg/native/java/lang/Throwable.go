package lang

import (
	"jvm/pkg/global"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func init() {
	native.MethodRegistry.Registry("java/lang/Throwable", "fillInStackTrace", "(I)Ljava/lang/Throwable;", fillInStackTrace)
}

func fillInStackTrace(frame *rtda.Frame) {
	thread := frame.Thread()
	method := frame.Method()
	loader := method.Class().ClassLoader()

	throwableObject := frame.LocalVars().GetThis().(*heap.NormalObject)

	StackTraceElementArrayClass := loader.LoadClass("[Ljava/lang/StackTraceElement;")
	StackTraceElementClass := StackTraceElementArrayClass.ElementClass()
	stackTrace := StackTraceElementArrayClass.NewArray(int32(thread.StackSize()))

	currentFrame := frame
	for i := int32(0); i < int32(thread.StackSize()) && currentFrame != nil; i++ {
		stackTrace.Set(i, newStackTraceElement(currentFrame, StackTraceElementClass))
		currentFrame = currentFrame.Lower()
	}

	throwableObject.SetExtend(stackTrace)

	frame.OperandStack().PushRef(throwableObject)
}

func newStackTraceElement(frame *rtda.Frame, StackTraceElementClass *heap.ClassObject) *heap.NormalObject {
	stackTraceObject := StackTraceElementClass.NewObject().(*heap.NormalObject)

	method := frame.Method()
	class := method.Class()
	loader := class.ClassLoader()

	StackTraceElementClass.SetField(stackTraceObject, "declaringClass", global.FdString, heap.JString(loader, class.JavaName()))
	StackTraceElementClass.SetField(stackTraceObject, "methodName", global.FdString, heap.JString(loader, method.Name()))
	StackTraceElementClass.SetField(stackTraceObject, "fileName", global.FdString, heap.JString(loader, class.SourceFile()))
	StackTraceElementClass.SetField(stackTraceObject, "lineNumber", global.FdInt, int32(method.GetLineNumber(frame.NextPC()-1)))

	return stackTraceObject
}
