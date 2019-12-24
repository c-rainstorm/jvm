package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type FALoad struct {
	base.NoOperandsInstruction
}

func (this *FALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushFloat(arrayRef.(*heap.ArrayObject).Get(index).(float32))
}
