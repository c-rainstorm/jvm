package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type AALoad struct {
	base.NoOperandsInstruction
}

func (this *AALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushRef(arrayRef.(*heap.ArrayObject).Get(index).(heap.Object))
}
