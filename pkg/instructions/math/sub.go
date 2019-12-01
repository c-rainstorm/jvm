package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// isub   Page.501
//
// 格式: isub
// 字节: 0x64
// 操作: 计算操作数栈顶两个整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2-value1)
// PS: value1 和 value2 都必须是整数，若发生溢出，则保留低 32位，溢出时不抛异常
//
type ISub struct {
	base.NoOperandsInstruction
}

func (this *ISub) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	val2 := operandStack.PopInt()
	operandStack.PushInt(val2 - val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lsub   Page.502
//
// 格式: lsub
// 字节: 0x65
// 操作: 计算操作数栈顶两个long型整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2-value1)
// PS: value1 和 value2 都必须是long，若发生溢出，则保留低 64位，溢出时不抛异常
//
type LSub struct {
	base.NoOperandsInstruction
}

func (this *LSub) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()
	operandStack.PushLong(val2 - val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fsub   Page.503
//
// 格式: fsub
// 字节: 0x66
// 操作: 计算操作数栈顶两个flaot的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2-value1)
// PS: value1 和 value2 都必须是float。特殊的计算规则（IEEE定义）
// 等同 value2 + (-value1)
type FSub struct {
	base.NoOperandsInstruction
}

func (this *FSub) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushFloat(val2 - val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dsub   Page.504
//
// 格式: dsub
// 字节: 0x67
// 操作: 计算操作数栈顶两个double的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2-value1)
// PS: value1 和 value2 都必须是double。特殊的计算规则（IEEE定义）
// 等同 value2 + (-value1)
//
type DSub struct {
	base.NoOperandsInstruction
}

func (this *DSub) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushDouble(val2 - val1)
}
