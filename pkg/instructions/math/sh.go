package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ishl
// 字节: 0x78
// 操作: 将栈顶整型元素左移 bits 位
//       ..., num, bits ->
//       ..., (num << (bits & 0x1F))
// PS: 因为是int，所以左移操作的位数不会超过 31，所以拿到 bits 后，截取后 5个bit 作为位移距离
//
type IShL struct {
	base.NoOperandsInstruction
}

func (this *IShL) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopInt()
	operandStack.PushInt(num << (uint32(bits) & 0x1F))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: ishr
// 字节: 0x7a
// 操作: 将栈顶整型元素右移 bits 位（算数右移）
//       ..., num, bits ->
//       ..., (num >> (bits & 0x1F))
// PS: 因为是int，所以位移操作的位数不会超过 31，所以拿到 bits 后，截取后 5个bit 作为位移距离
//
type IShR struct {
	base.NoOperandsInstruction
}

func (this *IShR) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopInt()
	operandStack.PushInt(num >> (uint32(bits) & 0x1F))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: iushr
// 字节: 0x7c
// 操作: 将栈顶整型元素逻辑右移 bits 位
//       ..., num, bits ->
//       ..., (num >> (bits & 0x1F))
// PS: 因为是int，所以位移操作的位数不会超过 31，所以拿到 bits 后，截取后 5个bit 作为位移距离
//
type IUShR struct {
	base.NoOperandsInstruction
}

func (this *IUShR) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopInt()
	operandStack.PushInt(int32(uint32(num) >> (uint32(bits) & 0x1F)))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lshl
// 字节: 0x79
// 操作: 将栈顶long型元素左移 bits 位
//       ..., num, bits ->
//       ..., (num << (bits & 0x3F))
// PS: num 是 long，bits是 int，左移操作的位数不会超过 63，所以拿到 bits 后，截取后 6个bit 作为位移距离
//
type LShL struct {
	base.NoOperandsInstruction
}

func (this *LShL) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopLong()
	operandStack.PushLong(num << (uint32(bits) & 0x3F))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lshr
// 字节: 0x7b
// 操作: 将栈顶long型元素右移 bits 位（算数右移）
//       ..., num, bits ->
//       ..., (num >> (bits & 0x3F))
// PS: num 是 long，bits是 int，左移操作的位数不会超过 63，所以拿到 bits 后，截取后 6个bit 作为位移距离
//
type LShR struct {
	base.NoOperandsInstruction
}

func (this *LShR) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopLong()
	operandStack.PushLong(num >> (uint32(bits) & 0x3F))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: lushr
// 字节: 0x7d
// 操作: 将栈顶long型元素逻辑右移 bits 位
//       ..., num, bits ->
//       ..., (num >> (bits & 0x3F))
// PS: num 是 long，bits是 int，左移操作的位数不会超过 63，所以拿到 bits 后，截取后 6个bit 作为位移距离
//
type LUShR struct {
	base.NoOperandsInstruction
}

func (this *LUShR) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	bits := operandStack.PopInt()
	num := operandStack.PopLong()
	operandStack.PushLong(int64(uint64(num) >> (uint32(bits) & 0x3F)))
}
