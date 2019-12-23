package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type ArrayLength struct {
	base.NoOperandsInstruction
}

func (this *ArrayLength) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	arrayRef := operandStack.PopRef()

	if arrayRef == nil {
		panic("java.lang.NullPointerException")
	}

	operandStack.PushInt(arrayRef.ArrayLength())
}
