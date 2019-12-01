package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: idiv
// 字节: 0x6c
// 操作: 计算操作数栈顶两个整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 />
//       ..., (value2/value1)
// PS: value1 和 value2 都必须是整数，不能整除时，向 0 取整
// 当除数为 0 时抛 ArithmeticException 异常
//
type IDiv struct {
	base.NoOperandsInstruction
}

func (this *IDiv) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	val2 := operandStack.PopInt()

	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushInt(val2 / val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ldiv
// 字节: 0x6d
// 操作: 计算操作数栈顶两个long型整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 />
//       ..., (value2/value1)
// PS: value1 和 value2 都必须是long，若发生溢出，则保留低 64位，溢出时不抛异常
type LDiv struct {
	base.NoOperandsInstruction
}

func (this *LDiv) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()

	if val1 == 0 {
		panic("java.lang.ArithmeticException: / by zero")
	}

	operandStack.PushLong(val2 / val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: fdiv
// 字节: 0x6e
// 操作: 计算操作数栈顶两个flaot的和，并将结果推到操作数栈顶
//       ..., value2, value1 />
//       ..., (value2/value1)
// PS: value1 和 value2 都必须是float。特殊的计算规则（IEEE定义）
//     value1 为 0 时，结果为无穷大
type FDiv struct {
	base.NoOperandsInstruction
}

func (this *FDiv) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushFloat(val2 / val1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ddiv
// 字节: 0x6f
// 操作: 计算操作数栈顶两个double的和，并将结果推到操作数栈顶
//       ..., value2, value1 />
//       ..., (value2/value1)
// PS: value1 和 value2 都必须是double。特殊的计算规则（IEEE定义）
//     value1 为 0 时，结果为无穷大
//
type DDiv struct {
	base.NoOperandsInstruction
}

func (this *DDiv) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushDouble(val2 / val1)
}
