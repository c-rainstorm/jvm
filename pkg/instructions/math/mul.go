package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: imul
// 字节: 0x68
// 操作: 计算操作数栈顶两个整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1*value2)
// PS: value1 和 value2 都必须是整数，若发生溢出，则保留低 32位，溢出时不抛异常
//
type IMul struct {
	base.NoOperandsInstruction
}

func (this *IMul) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	val2 := operandStack.PopInt()
	operandStack.PushInt(val1 * val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lmul
// 字节: 0x69
// 操作: 计算操作数栈顶两个long型整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1*value2)
// PS: value1 和 value2 都必须是long，若发生溢出，则保留低 64位，溢出时不抛异常
//
type LMul struct {
	base.NoOperandsInstruction
}

func (this *LMul) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()
	operandStack.PushLong(val1 * val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: fmul
// 字节: 0x6a
// 操作: 计算操作数栈顶两个flaot的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1*value2)
// PS: value1 和 value2 都必须是float。特殊的计算规则（IEEE定义）
// 1. value1 或 value2 为 NaN 的时候，结果是 NaN
// 2. 都不是NaN时，若两个操作数符号相同，则结果符号为正，否则为负
// 3. 无穷大 * 0.0 = NaN
// 4. 任何无穷大的数和非无穷大的数相乘，结果为无穷大，符号由规则2定义
// 5. 算数操作后可能会出现溢出（向上，向下），不抛异常（若 计算结果float无法表示，则会转成最相近的数）
// 术语定义：
//     Float 数格式：https://en.wikipedia.org/wiki/IEEE_754-1985#/media/File:IEEE_754_Single_Floating_Point_Format.svg
//     1. NaN: sign(1bit), exponent(8bits),fraction(23bit)
//          sign = either 0 or 1.
//          biased exponent = all 1 bits.
//          fraction = anything except all 0 bits (since all 0 bits represents infinity).
//
type FMul struct {
	base.NoOperandsInstruction
}

func (this *FMul) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushFloat(val1 * val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: dmul
// 字节: 0x6b
// 操作: 计算操作数栈顶两个double的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1*value2)
// PS: value1 和 value2 都必须是double。特殊的计算规则（IEEE定义）
// 1. value1 或 value2 为 NaN 的时候，结果是 NaN
// 2. 都不是NaN时，若两个操作数符号相同，则结果符号为正，否则为负
// 3. 无穷大 * 0.0 = NaN
// 4. 任何无穷大的数和非无穷大的数相乘，结果为无穷大，符号由规则2定义
// 5. 算数操作后可能会出现溢出（向上，向下），不抛异常（若 计算结果float无法表示，则会转成最相近的数）
// 术语定义：
//     Float 数格式：https://en.wikipedia.org/wiki/IEEE_754-1985#/media/File:IEEE_754_Single_Floating_Point_Format.svg
//     1. NaN: sign(1bit), exponent(8bits),fraction(23bit)
//          sign = either 0 or 1.
//          biased exponent = all 1 bits.
//          fraction = anything except all 0 bits (since all 0 bits represents infinity).
//
type DMul struct {
	base.NoOperandsInstruction
}

func (this *DMul) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushDouble(val1 * val2)
}
