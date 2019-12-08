package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type InvokeSpecial struct {
	base.Index16Instruction
}

func (this *InvokeSpecial) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopRef()
}
