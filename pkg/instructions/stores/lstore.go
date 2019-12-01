package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lstore   Page.533
//
// 格式: lstore [index]
// 字节: 0x37 0x01
//   指令样例     8进制表示     指令含义
//   lstore 1    0x37 0x01    将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 1 的位置
// 操作: 将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 index 的位置
// 描述: index 是一个 无符号的单字节整数
// 
//
type LStore struct {
	base.Index8Instruction
}

func (this *LStore) Execute(frame *rtda.Frame) {
	_LStore(frame, this.Index)
}

func _LStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopLong()
	frame.LocalVars().SetLong(index, val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lstore_<n>   Page.534
//
// 格式: lstore_<n>
// 字节:
//   指令样例     8进制表示    指令含义
//   lstore_0     0x3f       将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 0 的位置
//   lstore_1     0x40       将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 1 的位置
//   lstore_2     0x41       将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 2 的位置
//   lstore_3     0x42       将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 3 的位置
// 操作: 将操作数栈栈顶long型数据存储到当前帧（方法）本地变量表下标为 <n> 的位置
// 描述:
//
type LStore0 struct {
	base.NoOperandsInstruction
}

func (this *LStore0) Execute(frame *rtda.Frame) {
	_LStore(frame, 0)
}

type LStore1 struct {
	base.NoOperandsInstruction
}

func (this *LStore1) Execute(frame *rtda.Frame) {
	_LStore(frame, 1)
}

type LStore2 struct {
	base.NoOperandsInstruction
}

func (this *LStore2) Execute(frame *rtda.Frame) {
	_LStore(frame, 2)
}

type LStore3 struct {
	base.NoOperandsInstruction
}

func (this *LStore3) Execute(frame *rtda.Frame) {
	_LStore(frame, 3)
}
