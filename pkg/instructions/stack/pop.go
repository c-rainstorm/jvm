package stack

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// pop   Page.548
//
// 格式: pop
// 字节: 0x57
// 操作: 弹出操作数栈栈顶的一个槽（若只有一次指令调用，弹出一个栈顶元素，则只能用于非 long 和 double 的数据）
//
type POP struct {
	base.NoOperandsInstruction
}

func (this *POP) Execute(frame *rtda.Frame) {
	frame.OperandStack().PopSlot()
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// pop2   Page.549
//
// 格式: pop2
// 字节: 0x58
// 操作: 弹出操作数栈栈顶的两个槽（若栈顶是double或long，则弹出的是一个数据，若不是，则弹出两个数据）
//
type POP2 struct {
	base.NoOperandsInstruction
}

func (this *POP2) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	operandStack.PopSlot()
	operandStack.PopSlot()
}
