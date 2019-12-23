package stores

import (
	"jvm/pkg/rtda"
)

type LAStore struct {
	XAStore
}

func (this *LAStore) Execute(frame *rtda.Frame) {
	this.Execute0(frame, func(stack *rtda.OperandStack) interface{} {
		return stack.PopLong()
	})
}
