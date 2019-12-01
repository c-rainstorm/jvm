package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// iadd   Page.452
//
// 格式: iadd
// 字节: 0x60
// 操作: 计算操作数栈顶两个整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1+value2)
// PS: value1 和 value2 都必须是整数，若发生溢出，则保留低 32位，溢出时不抛异常
//
type IAdd struct {
	base.NoOperandsInstruction
}

func (this *IAdd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopInt()
	val2 := operandStack.PopInt()
	operandStack.PushInt(val1 + val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// ladd   Page.509
//
// 格式: ladd
// 字节: 0x61
// 操作: 计算操作数栈顶两个long型整数的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1+value2)
// PS: value1 和 value2 都必须是long，若发生溢出，则保留低 64位，溢出时不抛异常
//
type LAdd struct {
	base.NoOperandsInstruction
}

func (this *LAdd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopLong()
	val2 := operandStack.PopLong()
	operandStack.PushLong(val1 + val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fadd   Page.509
//
// 格式: fadd
// 字节: 0x62
// 操作: 计算操作数栈顶两个flaot的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1+value2)
// PS: value1 和 value2 都必须是float。特殊的计算规则（IEEE定义）
// 1. value1 或 value2 为 NaN 的时候，结果是 NaN
// 2. 相反符号无穷大相加 结果为 NaN  -
// 3. 相同符号无穷大相加 结果为 当下符号的无穷大
// 4. 任何无穷大的数和非无穷大的数相加，结果为该无穷大
// 5. -0.0 + +0.0 = +0.0
// 6. -0.0 + -0.0 = -0.0
// 7. +0.0 + +0.0 = +0.0
// 8. 0.0 + 非0.0 = 非0.0
// 9. -x + +x = +0.0
// 10. 算数操作后可能会出现溢出（向上，向下），不抛异常（若 计算结果float无法表示，则会转成最相近的数）
// 术语定义：
//     Float 数格式：https://en.wikipedia.org/wiki/IEEE_754-1985#/media/File:IEEE_754_Single_Floating_Point_Format.svg
//     1. NaN: sign(1bit), exponent(8bits),fraction(23bit)
//          sign = either 0 or 1.
//          biased exponent = all 1 bits.
//          fraction = anything except all 0 bits (since all 0 bits represents infinity).
//
type FAdd struct {
	base.NoOperandsInstruction
}

func (this *FAdd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopFloat()
	val2 := operandStack.PopFloat()
	operandStack.PushFloat(val1 + val2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dadd   Page.509
//
// 格式: dadd
// 字节: 0x63
// 操作: 计算操作数栈顶两个double的和，并将结果推到操作数栈顶
//       ..., value2, value1 ->
//       ..., (value1+value2)
// PS: value1 和 value2 都必须是double。特殊的计算规则（IEEE定义）
// 1. value1 或 value2 为 NaN 的时候，结果是 NaN
// 2. 相反符号无穷大相加 结果为 NaN  -
// 3. 相同符号无穷大相加 结果为 当下符号的无穷大
// 4. 任何无穷大的数和非无穷大的数相加，结果为该无穷大
// 5. -0.0 + +0.0 = +0.0
// 6. -0.0 + -0.0 = -0.0
// 7. +0.0 + +0.0 = +0.0
// 8. 0.0 + 非0.0 = 非0.0
// 9. -x + +x = +0.0
// 10. 算数操作后可能会出现溢出（向上，向下），不抛异常（若 计算结果double无法表示，则会转成最相近的数）
// 术语定义：
//     Float 数格式：https://en.wikipedia.org/wiki/IEEE_754-1985#/media/File:IEEE_754_Single_Floating_Point_Format.svg
//     1. NaN: sign(1bit), exponent(8bits),fraction(23bit)
//          sign = either 0 or 1.
//          biased exponent = all 1 bits.
//          fraction = anything except all 0 bits (since all 0 bits represents infinity).
//
type DAdd struct {
	base.NoOperandsInstruction
}

func (this *DAdd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val1 := operandStack.PopDouble()
	val2 := operandStack.PopDouble()
	operandStack.PushDouble(val1 + val2)
}
