package stores

import (
	"jvm/pkg/rtda"
)

type AAStore struct {
	XAStore
}

func (this *AAStore) Execute(frame *rtda.Frame) {
	this.Execute0(frame, func(stack *rtda.OperandStack) interface{} {
		return stack.PopRef()
	})
}
