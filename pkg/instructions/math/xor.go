package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ixor
// 字节: 0x82
// 操作: 栈顶两个 int 型操作数做异或操作
//       ..., num2, num1 ->
//       ..., (num2 ^ num1)
//
type IXOr struct {
	base.NoOperandsInstruction
}

func (this *IXOr) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopInt()
	num2 := operandStack.PopInt()
	operandStack.PushInt(num1 ^ num2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: land
// 字节: 0x83
// 操作: 栈顶两个 long 型操作数做异或操作
//       ..., num2, num1 ->
//       ..., (num2 ^ num1)
//
type LXOr struct {
	base.NoOperandsInstruction
}

func (this *LXOr) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopLong()
	num2 := operandStack.PopLong()
	operandStack.PushLong(num1 ^ num2)
}
