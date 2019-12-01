package math

import (
	"math"

	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: irem
// 字节: 0x70
// 操作: 计算操作数栈顶两个整数的余数，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2%value1)
// PS: value1 和 value2 都必须是整数，
// 当除数为 0 时抛 ArithmeticException 异常
//
type IRem struct {
	base.NoOperandsInstruction
}

func (this *IRem) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	val2 := operandStack.PopInt()

	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushInt(val2 % val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lrem
// 字节: 0x71
// 操作: 计算操作数栈顶两个long型整数的余数，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2%value1)
// PS: value1 和 value2 都必须是long
type LRem struct {
	base.NoOperandsInstruction
}

func (this *LRem) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()

	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushLong(val2 % val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: frem
// 字节: 0x72
// 操作: 计算操作数栈顶两个flaot的余数，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2%value1)
// PS: value1 和 value2 都必须是float。特殊的计算规则（IEEE定义）
type FRem struct {
	base.NoOperandsInstruction
}

func (this *FRem) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushFloat(float32(math.Mod(float64(val2), float64(val1))))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: drem
// 字节: 0x73
// 操作: 计算操作数栈顶两个double的余数，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value2 % value1)
// PS: value1 和 value2 都必须是double。特殊的计算规则（IEEE定义）
//
type DRem struct {
	base.NoOperandsInstruction
}

func (this *DRem) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushDouble(math.Mod(val2, val1))
}
