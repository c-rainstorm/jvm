package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// istore   Page.499
//
// 格式: istore [index]
// 字节: 0x36 0x01
//   指令样例     8进制表示     指令含义
//   istore 1    0x36 0x01    将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 1 的位置
// 操作: 将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 index 的位置
// 描述: index 是一个 无符号的单字节整数
// 
//
type IStore struct {
	base.Index8Instruction
}

func (this *IStore) Execute(frame *rtda.Frame) {
	_IStore(frame, this.Index)
}

func _IStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopInt()
	frame.LocalVars().SetInt(index, val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// istore_<n>   Page.500
//
// 格式: istore_<n>
// 字节:
//   指令样例     8进制表示    指令含义
//   istore_0     0x3b       将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 0 的位置
//   istore_1     0x3c       将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 1 的位置
//   istore_2     0x3d       将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 2 的位置
//   istore_3     0x3e       将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 3 的位置
// 操作: 将操作数栈栈顶int型数据存储到当前帧（方法）本地变量表下标为 <n> 的位置
// 描述:
//
type IStore0 struct {
	base.NoOperandsInstruction
}

func (this *IStore0) Execute(frame *rtda.Frame) {
	_IStore(frame, 0)
}

type IStore1 struct {
	base.NoOperandsInstruction
}

func (this *IStore1) Execute(frame *rtda.Frame) {
	_IStore(frame, 1)
}

type IStore2 struct {
	base.NoOperandsInstruction
}

func (this *IStore2) Execute(frame *rtda.Frame) {
	_IStore(frame, 2)
}

type IStore3 struct {
	base.NoOperandsInstruction
}

func (this *IStore3) Execute(frame *rtda.Frame) {
	_IStore(frame, 3)
}
