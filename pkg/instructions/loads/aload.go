package loads

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// aload   Page.371
//
// 格式: aload [index]
// 字节: 0x19 0x01
//   指令样例     8进制表示     指令含义
//   aload 1     0x19 0x01    将当前帧（方法）本地变量表下标为 1 的对象引用推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 index 下标处加载对象引用到操作数栈
// 描述: index 是一个 无符号的单字节整数
// PS:  aload 类指令无法加载 returnAddress 类型的值到操作数栈，和 astore 指令的这种不对称是有意设计的
//
type ALoad struct {
	base.Index8Instruction
}

func (this *ALoad) Execute(frame *rtda.Frame) {
	_ALoad(frame, uint(this.Index))
}

func _ALoad(frame *rtda.Frame, index uint) {
	val := frame.LocalVars().GetRef(index)
	frame.OperandStack().PushRef(val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// aload_<n>   Page.372
//
// 格式: aload_<n>
// 字节:
//   指令样例     8进制表示   指令含义
//   aload_0     0x2a       将当前帧（方法）本地变量表下标为 0 的对象引用推到操作数栈
//   aload_1     0x2b       将当前帧（方法）本地变量表下标为 1 的对象引用推到操作数栈
//   aload_2     0x2c       将当前帧（方法）本地变量表下标为 2 的对象引用推到操作数栈
//   aload_3     0x2d       将当前帧（方法）本地变量表下标为 3 的对象引用推到操作数栈
// 操作: 从当前帧（方法）本地变量表的 <n> 下标处加载对象引用到操作数栈
// 描述:
//
type ALoad0 struct {
	base.NoOperandsInstruction
}

func (this *ALoad0) Execute(frame *rtda.Frame) {
	_ALoad(frame, 0)
}

type ALoad1 struct {
	base.NoOperandsInstruction
}

func (this *ALoad1) Execute(frame *rtda.Frame) {
	_ALoad(frame, 1)
}

type ALoad2 struct {
	base.NoOperandsInstruction
}

func (this *ALoad2) Execute(frame *rtda.Frame) {
	_ALoad(frame, 2)
}

type ALoad3 struct {
	base.NoOperandsInstruction
}

func (this *ALoad3) Execute(frame *rtda.Frame) {
	_ALoad(frame, 3)
}
