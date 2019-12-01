package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lload   Page.521
//
// 格式: lload [index]
// 字节: 0x16 0x01
//   指令样例     8进制表示     指令含义
//   lload 1     0x16 0x01    将当前帧（方法）本地变量表下标为 1 的long型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 index 下标处加载long型变量到操作数栈
// 描述: index 是一个 无符号的单字节整数
//
type LLoad struct {
	base.Index8Instruction
}

func (this *LLoad) Execute(frame *rtda.Frame) {
	_LLoad(frame, this.Index)
}

func _LLoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetLong(index)
	frame.OperandStack().PushLong(val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lload_<n>   Page.522
//
// 格式: lload_<n>
// 字节:
//   指令样例     8进制表示   指令含义
//   lload_0     0x1e       将当前帧（方法）本地变量表下标为 0 的long型变量推到操作数栈
//   lload_1     0x1f       将当前帧（方法）本地变量表下标为 1 的long型变量推到操作数栈
//   lload_2     0x20       将当前帧（方法）本地变量表下标为 2 的long型变量推到操作数栈
//   lload_3     0x21       将当前帧（方法）本地变量表下标为 3 的long型变量推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 <n> 下标处加载long型变量到操作数栈, n, n+1 都必须是本地变量表的合法下标
// 描述:
//
type LLoad0 struct {
	base.NoOperandsInstruction
}

func (this *LLoad0) Execute(frame *rtda.Frame) {
	_LLoad(frame, 0)
}

type LLoad1 struct {
	base.NoOperandsInstruction
}

func (this *LLoad1) Execute(frame *rtda.Frame) {
	_LLoad(frame, 1)
}

type LLoad2 struct {
	base.NoOperandsInstruction
}

func (this *LLoad2) Execute(frame *rtda.Frame) {
	_LLoad(frame, 2)
}

type LLoad3 struct {
	base.NoOperandsInstruction
}

func (this *LLoad3) Execute(frame *rtda.Frame) {
	_LLoad(frame, 3)
}
