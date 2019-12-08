package references

import (
	"fmt"

	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type InvokeVirtual struct {
	base.Index16Instruction
}

func (this *InvokeVirtual) Execute(frame *rtda.Frame) {
	methodRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.MethodSymRef)
	if methodRef.Name() == "println" {
		operandStack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", operandStack.PopInt() != 0)
		case "(B)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(C)V":
			fmt.Printf("%c\n", operandStack.PopInt())
		case "(S)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(I)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(J)V":
			fmt.Printf("%v\n", operandStack.PopLong())
		case "(F)V":
			fmt.Printf("%v\n", operandStack.PopFloat())
		case "(D)V":
			fmt.Printf("%v\n", operandStack.PopDouble())
		default:
			panic("println: " + methodRef.Descriptor())
		}

		frame.OperandStack().PopRef()
	}
}
