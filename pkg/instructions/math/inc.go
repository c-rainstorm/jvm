package math

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: iinc [index] [const] (index, const 都是一字节)
// 字节: 0x84 0x01 0x02  // 给本地变量表下标为 1 的变量加上常量 2
// 操作: 给本地变量表下标为 index 的整型变量加上常量 const
//
type IInc struct {
	base.Index8Instruction
	Const int32
}

func (this *IInc) FetchOperands(reader *base.ByteCodeReader) {
	this.Index8Instruction.FetchOperands(reader)
	this.Const = int32(reader.ReadUint8())
}

func (this *IInc) Execute(frame *rtda.Frame) {
	localVars := frame.LocalVars()
	index := uint(this.Index)
	val := localVars.GetInt(index)
	localVars.SetInt(index, val+this.Const)
}
