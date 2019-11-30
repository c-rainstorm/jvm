package stack

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// swap   Page.559
//
// 格式: swap
// 字节: 0x5f
// 操作: ..., slot2, slot1 ->
//       ..., slot1, slot2
// PS: 交换栈顶的两个槽，目前版本的 JVM 不提供 long、double 类型操作数的 swap 操作
//
type Swap struct {
	base.NoOperandsInstruction
}

func (this *Swap) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot2)
}
