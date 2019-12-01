package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dload   Page.399
//
// 格式: dload [index]
// 字节: 0x18 0x01
//   指令样例     8进制表示     指令含义
//   dload 1     0x18 0x01    将当前帧（方法）本地变量表下标为 1 的double型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 index 下标处加载double型变量到操作数栈
// 描述: index 是一个 无符号的单字节整数
//
type DLoad struct {
	base.Index8Instruction
}

func (this *DLoad) Execute(frame *rtda.Frame) {
	_DLoad(frame, this.Index)
}

func _DLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetDouble(index)
	frame.OperandStack().PushDouble(val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dload_<n>   Page.400
//
// 格式: dload_<n>
// 字节:
//   指令样例     8进制表示   指令含义
//   dload_0     0x1e       将当前帧（方法）本地变量表下标为 0 的double型变量推到操作数栈
//   dload_1     0x1f       将当前帧（方法）本地变量表下标为 1 的double型变量推到操作数栈
//   dload_2     0x20       将当前帧（方法）本地变量表下标为 2 的double型变量推到操作数栈
//   dload_3     0x21       将当前帧（方法）本地变量表下标为 3 的double型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 <n> 下标处加载double型变量到操作数栈, n, n+1 都必须是本地变量表的合法下标
// 描述:
//
type DLoad0 struct {
	base.NoOperandsInstruction
}

func (this *DLoad0) Execute(frame *rtda.Frame) {
	_DLoad(frame, 0)
}

type DLoad1 struct {
	base.NoOperandsInstruction
}

func (this *DLoad1) Execute(frame *rtda.Frame) {
	_DLoad(frame, 1)
}

type DLoad2 struct {
	base.NoOperandsInstruction
}

func (this *DLoad2) Execute(frame *rtda.Frame) {
	_DLoad(frame, 2)
}

type DLoad3 struct {
	base.NoOperandsInstruction
}

func (this *DLoad3) Execute(frame *rtda.Frame) {
	_DLoad(frame, 3)
}
