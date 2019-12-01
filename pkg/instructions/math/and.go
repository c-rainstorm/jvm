package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: iand
// 字节: 0x7e
// 操作: 栈顶两个 int 型操作数做与操作
//       ..., num2, num1 ->
//       ..., (num2 & num1)
//
type IAnd struct {
	base.NoOperandsInstruction
}

func (this *IAnd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopInt()
	num2 := operandStack.PopInt()
	operandStack.PushInt(num1 & num2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: land
// 字节: 0x7f
// 操作: 栈顶两个 long 型操作数做与操作
//       ..., num2, num1 ->
//       ..., (num2 & num1)
//
type LAnd struct {
	base.NoOperandsInstruction
}

func (this *LAnd) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopLong()
	num2 := operandStack.PopLong()
	operandStack.PushLong(num1 & num2)
}
