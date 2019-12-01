package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ineg
// 字节: 0x74
// 操作: 弹出操作数栈顶的整数取反后，并将结果推到操作数栈顶
//       ..., value1 ->
//       ..., (-value1)
//
type INeg struct {
	base.NoOperandsInstruction
}

func (this *INeg) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	operandStack.PushInt(-val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lneg
// 字节: 0x75
// 操作: 弹出操作数栈顶的long取反后，并将结果推到操作数栈顶
//       ..., value1 ->
//       ..., (-value1)
type LNeg struct {
	base.NoOperandsInstruction
}

func (this *LNeg) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	operandStack.PushLong(-val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: fneg
// 字节: 0x76
// 操作: 弹出操作数栈顶的float取反后，并将结果推到操作数栈顶
//       ..., value1 ->
//       ..., (-value1)
type FNeg struct {
	base.NoOperandsInstruction
}

func (this *FNeg) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	operandStack.PushFloat(-val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: dneg
// 字节: 0x77
// 操作: 弹出操作数栈顶的double取反后，并将结果推到操作数栈顶
//       ..., value1 ->
//       ..., (-value1)
//
type DNeg struct {
	base.NoOperandsInstruction
}

func (this *DNeg) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	operandStack.PushDouble(-val1)
}
