package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type LALoad struct {
	base.NoOperandsInstruction
}

func (this *LALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushLong(arrayRef.Get(index).(int64))
}
