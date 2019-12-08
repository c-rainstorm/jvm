package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type InstanceOf struct {
	base.Index16Instruction
}

func (this *InstanceOf) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	ref := operandStack.PopRef()
	if ref == nil {
		operandStack.PushInt(0)
		return
	}

	classRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.ClassSymRef)
	class := classRef.ResolvedClass()

	if ref.IsInstanceOf(class) {
		operandStack.PushInt(1)
	} else {
		operandStack.PushInt(0)
	}
}
