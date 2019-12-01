package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

//  ..., value2, value1 ->
//  ..., (value2-value1的符号)
type LCmp struct {
	base.NoOperandsInstruction
}

func (this *LCmp) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()

	operandStack.PushInt(SignLong(val2 - val1))
}
