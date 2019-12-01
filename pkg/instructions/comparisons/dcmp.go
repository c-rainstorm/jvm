package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

//  ..., value2, value1 ->
//  ..., (value2-value1的符号)，当 其中有不能比较的数值时，根据指令后缀的 G/L 决定比较结果是 1/-1
type DCmpG struct {
	base.NoOperandsInstruction
}

func (this *DCmpG) Execute(frame *rtda.Frame) {
	DCmp_(frame, true)
}

type DCmpL struct {
	base.NoOperandsInstruction
}

func (this *DCmpL) Execute(frame *rtda.Frame) {
	DCmp_(frame, false)
}

func DCmp_(frame *rtda.Frame, gFlag bool) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushInt(SignDouble(val2-val1, gFlag))
}
