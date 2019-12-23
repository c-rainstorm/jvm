package stores

import (
	"jvm/pkg/rtda"
)

type DAStore struct {
	XAStore
}

func (this *DAStore) Execute(frame *rtda.Frame) {
	this.Execute0(frame, func(stack *rtda.OperandStack) interface{} {
		return stack.PopDouble()
	})
}
