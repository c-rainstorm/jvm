package references

import (
	"fmt"

	"jvm/pkg/global"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type AThrow struct {
	base.NoOperandsInstruction
}

func (this *AThrow) Execute(frame *rtda.Frame) {
	exception := frame.OperandStack().PopRef()
	if exception == nil {
		panic("java.lang.NullPointerException")
	}

	thread := frame.Thread()

	handler := findExceptionHandler(thread, exception.Class())
	if handler == nil {
		printStackTrace(exception.(*heap.NormalObject))
		return
	}

	executeHandler(thread, exception, handler)
}

func executeHandler(thread *rtda.Thread, exception heap.Object, handler *heap.ExceptionHandler) {
	frame := thread.TopFrame()
	frame.OperandStack().Clear()
	frame.OperandStack().PushRef(exception)

	frame.SetNextPC(int(handler.HandlerPc()))
}

func printStackTrace(throwableObject *heap.NormalObject) {
	stackTrace := throwableObject.GetExtend().(*heap.ArrayObject)
	arrayLength := stackTrace.ArrayLength()
	if arrayLength == 0 {
		return
	}

	StackTraceElementClass := stackTrace.Class().ElementClass()
	detailMessage := heap.GoString(throwableObject.Class().GetField(throwableObject, "detailMessage", global.FdString).(*heap.NormalObject))
	fmt.Printf("%s: %s\n", throwableObject.Class().JavaName(), detailMessage)

	for i := int32(0); i < arrayLength; i++ {
		printStackTrace0(stackTrace.Get(i).(*heap.NormalObject), StackTraceElementClass)
	}
}

func printStackTrace0(stackTraceObject *heap.NormalObject, StackTraceElementClass *heap.ClassObject) {
	methodName := heap.GoString(StackTraceElementClass.GetField(stackTraceObject, "methodName", global.FdString).(*heap.NormalObject))
	if (methodName != "fillInStackTrace") && (methodName != "<init>") {
		classname := heap.GoString(StackTraceElementClass.GetField(stackTraceObject, "declaringClass", global.FdString).(*heap.NormalObject))
		fileName := heap.GoString(StackTraceElementClass.GetField(stackTraceObject, "fileName", global.FdString).(*heap.NormalObject))
		lineNumber := StackTraceElementClass.GetField(stackTraceObject, "lineNumber", global.FdInt).(int32)
		// at java.lang.NumberFormatException.forInputString(NumberFormatException.java:65)
		fmt.Printf("\tat %s.%s(%s:%d)\n", classname, methodName, fileName, lineNumber)
	}
}

func findExceptionHandler(thread *rtda.Thread, exceptionClass *heap.ClassObject) *heap.ExceptionHandler {
	for !thread.IsStackEmpty() {
		currentFrame := thread.TopFrame()
		method := currentFrame.Method()

		handler := method.LookupExceptionTable(exceptionClass, uint16(currentFrame.NextPC()-1))
		if handler != nil {
			return handler
		}
		thread.PopFrame()
	}
	return nil
}
