package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type DALoad struct {
	base.NoOperandsInstruction
}

func (this *DALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushDouble(arrayRef.Get(index).(float64))
}
