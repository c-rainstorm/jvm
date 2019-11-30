package constants

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// nop   Page.547
//
// 格式: nop
// 字节: 0x0
// 操作: do nothing
//
type NOP struct {
	base.NoOperandsInstruction
}

func (this *NOP) Execute(frame *rtda.Frame) {
	// do nothing
}