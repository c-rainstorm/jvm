package conversions

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: l2i
// 字节: 0x88
// 操作: 将栈顶的 long 型转成 int 型
//
type L2I struct {
	base.NoOperandsInstruction
}

func (this *L2I) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	longVal := operandStack.PopLong()
	operandStack.PushInt(int32(longVal))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: l2f
// 字节: 0x89
// 操作: 将栈顶的 long 型转成 float 型
//
type L2F struct {
	base.NoOperandsInstruction
}

func (this *L2F) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopLong()
	operandStack.PushFloat(float32(val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: l2d
// 字节: 0x8a
// 操作: 将栈顶的 long 型转成 double 型
//
type L2D struct {
	base.NoOperandsInstruction
}

func (this *L2D) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopLong()
	operandStack.PushDouble(float64(val))
}
