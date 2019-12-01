package conversions

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: f2i
// 字节: 0x8b
// 操作: 将栈顶的 float 型转成 int 型
//
type F2I struct {
	base.NoOperandsInstruction
}

func (this *F2I) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopFloat()
	operandStack.PushInt(int32(val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: f2l
// 字节: 0x8c
// 操作: 将栈顶的 float 型转成 long 型
//
type F2L struct {
	base.NoOperandsInstruction
}

func (this *F2L) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopFloat()
	operandStack.PushLong(int64(val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: f2d
// 字节: 0x8d
// 操作: 将栈顶的 float 型转成 double 型
//
type F2D struct {
	base.NoOperandsInstruction
}

func (this *F2D) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	val := operandStack.PopFloat()
	operandStack.PushDouble(float64(val))
}
