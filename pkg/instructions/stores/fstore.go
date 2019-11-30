package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fstore   Page.437
//
// 格式: fstore [index]
// 字节: ox38 0x01
//   指令样例     8进制表示     指令含义
//   fstore 1    ox38 0x01    将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 1 的位置
// 操作: 将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 index 的位置
// 描述: index 是一个 无符号的单字节整数
// 
//
type FStore struct {
	base.Index8Instruction
}

func (this *FStore) Execute(frame *rtda.Frame) {
	_FStore(frame, uint(this.Index))
}

func _FStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopFloat()
	frame.LocalVars().SetFloat(index, val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fstore_<n>   Page.438
//
// 格式: fstore_<n>
// 字节:
//   指令样例     8进制表示    指令含义
//   fstore_0     0x3b       将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 0 的位置
//   fstore_1     0x3c       将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 1 的位置
//   fstore_2     0x3d       将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 2 的位置
//   fstore_3     0x3e       将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 3 的位置
// 操作: 将操作数栈栈顶float型数据存储到当前帧（方法）本地变量表下标为 <n> 的位置
// 描述:
//
type FStore0 struct {
	base.NoOperandsInstruction
}

func (this *FStore0) Execute(frame *rtda.Frame) {
	_FStore(frame, 0)
}

type FStore1 struct {
	base.NoOperandsInstruction
}

func (this *FStore1) Execute(frame *rtda.Frame) {
	_FStore(frame, 1)
}

type FStore2 struct {
	base.NoOperandsInstruction
}

func (this *FStore2) Execute(frame *rtda.Frame) {
	_FStore(frame, 2)
}

type FStore3 struct {
	base.NoOperandsInstruction
}

func (this *FStore3) Execute(frame *rtda.Frame) {
	_FStore(frame, 3)
}
