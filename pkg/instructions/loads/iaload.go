package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type IALoad struct {
	base.NoOperandsInstruction
}

func (this *IALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushInt(arrayRef.(*heap.ArrayObject).Get(index).(int32))
}

func NotNull(ref heap.Object) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
