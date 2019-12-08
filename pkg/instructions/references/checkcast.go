package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type CheckCast struct {
	base.Index16Instruction
}

func (this *CheckCast) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	ref := operandStack.PopRef()
	operandStack.PushRef(ref)
	if ref == nil {
		return
	}

	classRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.ClassSymRef)
	class := classRef.ResolvedClass()

	if !ref.IsInstanceOf(class) {
		panic("java.lang.ClassCastException")
	}
}
