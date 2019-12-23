package stores

import (
	"jvm/pkg/rtda"
)

type IAStore struct {
	XAStore
}

func (this *IAStore) Execute(frame *rtda.Frame) {
	this.Execute0(frame, func(stack *rtda.OperandStack) interface{} {
		return stack.PopInt()
	})
}