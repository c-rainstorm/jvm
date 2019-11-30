package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// iload   Page.466
//
// 格式: iload [index]
// 字节: 0x15 0x01
//   指令样例     8进制表示     指令含义
//   iload 1     0x15 0x01    将当前帧（方法）本地变量表下标为 1 的整型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 index 下标处加载整型变量到操作数栈
// 描述: index 是一个 无符号的单字节整数
//
type ILoad struct {
	base.Index8Instruction
}

func (this *ILoad) Execute(frame *rtda.Frame) {
	_ILoad(frame, uint(this.Index))
}

func _ILoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetInt(index)
	frame.OperandStack().PushInt(val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// iload_<n>   Page.467
//
// 格式: iload_<n>
// 字节:
//   指令样例     8进制表示   指令含义
//   iload_0     0x1a       将当前帧（方法）本地变量表下标为 0 的整型变量推到操作数栈
//   iload_1     0x1b       将当前帧（方法）本地变量表下标为 1 的整型变量推到操作数栈
//   iload_2     0x1c       将当前帧（方法）本地变量表下标为 2 的整型变量推到操作数栈
//   iload_3     0x1d       将当前帧（方法）本地变量表下标为 3 的整型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 <n> 下标处加载整型变量到操作数栈
// 描述:
//
type ILoad0 struct {
	base.NoOperandsInstruction
}

func (this *ILoad0) Execute(frame *rtda.Frame) {
	_ILoad(frame, 0)
}

type ILoad1 struct {
	base.NoOperandsInstruction
}

func (this *ILoad1) Execute(frame *rtda.Frame) {
	_ILoad(frame, 1)
}

type ILoad2 struct {
	base.NoOperandsInstruction
}

func (this *ILoad2) Execute(frame *rtda.Frame) {
	_ILoad(frame, 2)
}

type ILoad3 struct {
	base.NoOperandsInstruction
}

func (this *ILoad3) Execute(frame *rtda.Frame) {
	_ILoad(frame, 3)
}
