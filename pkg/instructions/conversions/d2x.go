package conversions

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: d2i
// 字节: 0x8e
// 操作: 将栈顶的 double 型转成 int 型
//
type D2I struct {
	base.NoOperandsInstruction
}

func (this *D2I) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopDouble()
	operandStack.PushInt(int32(val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: d2l
// 字节: 0x8f
// 操作: 将栈顶的 double 型转成 long 型
//
type D2L struct {
	base.NoOperandsInstruction
}

func (this *D2L) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopDouble()
	operandStack.PushLong(int64(val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: d2f
// 字节: 0x90
// 操作: 将栈顶的 double 型转成 float 型
//
type D2F struct {
	base.NoOperandsInstruction
}

func (this *D2F) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopDouble()
	operandStack.PushFloat(float32(val))
}
