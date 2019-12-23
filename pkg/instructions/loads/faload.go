package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type FALoad struct {
	base.NoOperandsInstruction
}

func (this *FALoad) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	operandStack.PushFloat(arrayRef.Get(index).(float32))
}
