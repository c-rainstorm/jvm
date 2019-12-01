package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fload   Page.429
//
// 格式: fload [index]
// 字节: 0x17 0x01
//   指令样例     8进制表示     指令含义
//   fload 1     0x17 0x01    将当前帧（方法）本地变量表下标为 1 的float型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 index 下标处加载float型变量到操作数栈
// 描述: index 是一个 无符号的单字节整数
//
type FLoad struct {
	base.Index8Instruction
}

func (this *FLoad) Execute(frame *rtda.Frame) {
	_FLoad(frame, this.Index)
}

func _FLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetFloat(index)
	frame.OperandStack().PushFloat(val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fload_<n>   Page.430
//
// 格式: fload_<n>
// 字节:
//   指令样例     8进制表示   指令含义
//   fload_0     0x22       将当前帧（方法）本地变量表下标为 0 的float型变量推到操作数栈
//   fload_1     0x23       将当前帧（方法）本地变量表下标为 1 的float型变量推到操作数栈
//   fload_2     0x24       将当前帧（方法）本地变量表下标为 2 的float型变量推到操作数栈
//   fload_3     0x25       将当前帧（方法）本地变量表下标为 3 的float型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 <n> 下标处加载float型变量到操作数栈
// 描述:
//
type FLoad0 struct {
	base.NoOperandsInstruction
}

func (this *FLoad0) Execute(frame *rtda.Frame) {
	_FLoad(frame, 0)
}

type FLoad1 struct {
	base.NoOperandsInstruction
}

func (this *FLoad1) Execute(frame *rtda.Frame) {
	_FLoad(frame, 1)
}

type FLoad2 struct {
	base.NoOperandsInstruction
}

func (this *FLoad2) Execute(frame *rtda.Frame) {
	_FLoad(frame, 2)
}

type FLoad3 struct {
	base.NoOperandsInstruction
}

func (this *FLoad3) Execute(frame *rtda.Frame) {
	_FLoad(frame, 3)
}
