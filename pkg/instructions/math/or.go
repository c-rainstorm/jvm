package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ior
// 字节: 0x80
// 操作: 栈顶两个 long 型操作数做或操作
//       ..., num2, num1 ->
//       ..., (num2 | num1)
//
type IOr struct {
	base.NoOperandsInstruction
}

func (this *IOr) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopInt()
	num2 := operandStack.PopInt()
	operandStack.PushInt(num1 | num2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lor
// 字节: 0x81
// 操作: 栈顶两个 long 型操作数做或操作
//       ..., num2, num1 ->
//       ..., (num2 | num1)
//
type LOr struct {
	base.NoOperandsInstruction
}

func (this *LOr) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	num1 := operandStack.PopLong()
	num2 := operandStack.PopLong()
	operandStack.PushLong(num1 | num2)
}
