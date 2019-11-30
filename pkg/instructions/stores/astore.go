package stores

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// astore   Page.407
//
// 格式: astore [index]
// 字节: ox3a 0x01
//   指令样例     8进制表示     指令含义
//   astore 1    ox3a 0x01    将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 1 的位置
// 操作: 将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 index 的位置
// 描述: index 是一个 无符号的单字节整数
// PS : 该引用类型可以是 returnAddress 或者普通对象引用
//
type AStore struct {
	base.Index8Instruction
}

func (this *AStore) Execute(frame *rtda.Frame) {
	_AStore(frame, uint(this.Index))
}

func _AStore(frame *rtda.Frame, index uint) {
	val := frame.OperandStack().PopRef()
	frame.LocalVars().SetRef(index, val)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// astore_<n>   Page.408
//
// 格式: astore_<n>
// 字节:
//   指令样例     8进制表示    指令含义
//   astore_0     0x47       将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 0 的位置
//   astore_1     0x48       将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 1 的位置
//   astore_2     0x49       将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 2 的位置
//   astore_3     0x50       将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 3 的位置
// 操作: 将操作数栈栈顶对象引用存储到当前帧（方法）本地变量表下标为 <n> 的位置
// 描述:
//
type AStore0 struct {
	base.NoOperandsInstruction
}

func (this *AStore0) Execute(frame *rtda.Frame) {
	_AStore(frame, 0)
}

type AStore1 struct {
	base.NoOperandsInstruction
}

func (this *AStore1) Execute(frame *rtda.Frame) {
	_AStore(frame, 1)
}

type AStore2 struct {
	base.NoOperandsInstruction
}

func (this *AStore2) Execute(frame *rtda.Frame) {
	_AStore(frame, 2)
}

type AStore3 struct {
	base.NoOperandsInstruction
}

func (this *AStore3) Execute(frame *rtda.Frame) {
	_AStore(frame, 3)
}
