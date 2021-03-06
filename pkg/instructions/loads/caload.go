package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type CALoad struct {
	base.NoOperandsInstruction
}

func (this *CALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushInt(int32(arrayRef.(*heap.ArrayObject).Get(index).(uint16)))
}
