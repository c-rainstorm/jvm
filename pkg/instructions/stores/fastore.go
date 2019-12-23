package stores

import (
	"jvm/pkg/rtda"
)

type FAStore struct {
	XAStore
}

func (this *FAStore) Execute(frame *rtda.Frame) {
	this.Execute0(frame, func(stack *rtda.OperandStack) interface{} {
		return stack.PopFloat()
	})
}
