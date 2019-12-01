package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

//  ..., value2, value1 ->
//  ..., (value2-value1的符号)，当 其中有不能比较的数值时，根据指令后缀的 G/L 决定比较结果是 1/-1
type FCmpG struct {
	base.NoOperandsInstruction
}

func (this *FCmpG) Execute(frame *rtda.Frame) {
	FCmp_(frame, true)
}

type FCmpL struct {
	base.NoOperandsInstruction
}

func (this *FCmpL) Execute(frame *rtda.Frame) {
	FCmp_(frame, false)
}

func FCmp_(frame *rtda.Frame, gFlag bool) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushInt(SignFloat(val2-val1, gFlag))
}
