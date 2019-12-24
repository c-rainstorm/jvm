package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type XAStore struct {
	base.NoOperandsInstruction
}

func (this *XAStore) Execute0(frame *rtda.Frame, popFunc func(stack *rtda.OperandStack) interface{}) {
	operandStack := frame.OperandStack()

	value := popFunc(operandStack)
	index := operandStack.PopInt()
	arrayRef := operandStack.PopRef()

	NotNull(arrayRef)

	arrayRef.(*heap.ArrayObject).Set(index, value)
}

func NotNull(ref interface{}) {
	if ref == nil {
		panic("java.lang.NullPointerException")
	}
}
