package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type ANewArray struct {
	base.Index16Instruction
}

func (this *ANewArray) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	classRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.ClassSymRef)
	elementClass := classRef.ResolvedClass()

	count := operandStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrayClass := elementClass.ArrayClass()

	operandStack.PushRef(arrayClass.NewArray(count))
}
