package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dstore   Page.407
//
// 格式: dstore [index]
// 字节: ox39 0x01
//   指令样例     8进制表示     指令含义
//   dstore 1    ox39 0x01    将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 1 的位置
// 操作: 将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 index 的位置
// 描述: index 是一个 无符号的单字节整数
//
type DStore struct {
	base.Index8Instruction
}

func (this *DStore) Execute(frame *rtda.Frame) {
	_DStore(frame, uint(this.Index))
}

func _DStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopDouble()
	frame.LocalVars().SetDouble(index, val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dstore_<n>   Page.408
//
// 格式: dstore_<n>
// 字节:
//   指令样例     8进制表示    指令含义
//   dstore_0     0x47       将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 0 的位置
//   dstore_1     0x48       将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 1 的位置
//   dstore_2     0x49       将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 2 的位置
//   dstore_3     0x50       将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 3 的位置
// 操作: 将操作数栈栈顶double型数据存储到当前帧（方法）本地变量表下标为 <n> 的位置
// 描述:
//
type DStore0 struct {
	base.NoOperandsInstruction
}

func (this *DStore0) Execute(frame *rtda.Frame) {
	_DStore(frame, 0)
}

type DStore1 struct {
	base.NoOperandsInstruction
}

func (this *DStore1) Execute(frame *rtda.Frame) {
	_DStore(frame, 1)
}

type DStore2 struct {
	base.NoOperandsInstruction
}

func (this *DStore2) Execute(frame *rtda.Frame) {
	_DStore(frame, 2)
}

type DStore3 struct {
	base.NoOperandsInstruction
}

func (this *DStore3) Execute(frame *rtda.Frame) {
	_DStore(frame, 3)
}
