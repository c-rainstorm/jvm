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
	itemRef := arrayRef.(*heap.ArrayObject).Get(index)
	if itemRef == nil {
		operandStack.PushRef(nil)
	} else {
		operandStack.PushRef(itemRef.(heap.Object))
	}
}
