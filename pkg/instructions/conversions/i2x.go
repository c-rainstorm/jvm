package conversions

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2l
// 字节: 0x85
// 操作: 将栈顶的 int 型转成 long 型
//
type I2L struct {
	base.NoOperandsInstruction
}

func (this *I2L) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushLong(int64(intVal))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2f
// 字节: 0x86
// 操作: 将栈顶的 int 型转成 float 型
//
type I2F struct {
	base.NoOperandsInstruction
}

func (this *I2F) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushFloat(float32(intVal))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2d
// 字节: 0x87
// 操作: 将栈顶的 int 型转成 double 型
//
type I2D struct {
	base.NoOperandsInstruction
}

func (this *I2D) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushDouble(float64(intVal))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2b
// 字节: 0x91                                   符号拓展
// 操作: 将栈顶的 int 型转成 byte 型     int -> byte -> int
//
type I2B struct {
	base.NoOperandsInstruction
}

func (this *I2B) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushInt(int32(int8(intVal)))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2c
// 字节: 0x92                               保留低16位
// 操作: 将栈顶的 int 型转成 char 型     int -> char -> int
//
type I2C struct {
	base.NoOperandsInstruction
}

func (this *I2C) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushInt(intVal & 0xFFFF)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: i2s
// 字节: 0x93                                     符号拓展
// 操作: 将栈顶的 int 型转成 short 型     int -> short -> int
//
type I2S struct {
	base.NoOperandsInstruction
}

func (this *I2S) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	intVal := operandStack.PopInt()
	operandStack.PushInt(int32(int16(intVal)))
}
