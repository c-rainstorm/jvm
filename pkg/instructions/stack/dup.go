package stack

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup   Page.410
//
// 格式: dup
// 字节: 0x59
// 操作: 复制操作数栈栈顶的一个槽，只能用于站一个槽的数据（非long、double 的数据）
//
type Dup struct {
	base.NoOperandsInstruction
}

func (this *Dup) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot := operandStack.PopSlot()
	operandStack.PushSlot(slot)
	operandStack.PushSlot(slot)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup_x1   Page.411
//
// 格式: dup_x1
// 字节: 0x5a
// 操作:  ..., value2, value1 ->
//       ..., value1, value2, value1
// PS: value2 不能是 long、double
//
type DupX1 struct {
	base.NoOperandsInstruction
}

func (this *DupX1) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	value1 := operandStack.PopSlot()
	value2 := operandStack.PopSlot()
	operandStack.PushSlot(value1)
	operandStack.PushSlot(value2)
	operandStack.PushSlot(value1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup_x2   Page.412
//
// 格式: dup_x2
// 字节: 0x5b
// 操作:  ..., slot3, slot2, slot1 ->
//       ..., slot1, slot33, slot2, slot1 ->
// PS: slot3 和 slot2 可以是一个操作数（double、long），也可以是两个
//
type DupX2 struct {
	base.NoOperandsInstruction
}

func (this *DupX2) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	slot3 := operandStack.PopSlot()
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot3)
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup2   Page.413
//
// 格式: dup2
// 字节: 0x5c
// 操作: ..., slot2, slot1 ->
//       ..., slot2, slot1, slot2, slot1
//
type Dup2 struct {
	base.NoOperandsInstruction
}

func (this *Dup2) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup2_x1   Page.414
//
// 格式: dup2_x1
// 字节: 0x5d
// 操作:  ...,slot3, slot2, slot1 ->
//        ..., slot2, slot1, slot3, slot2, slot1
// PS: slot3 不能是 long、double 的其中一个槽
//
type Dup2X1 struct {
	base.NoOperandsInstruction
}

func (this *Dup2X1) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	slot3 := operandStack.PopSlot()
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot3)
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dup2_x2   Page.415
//
// 格式: dup2_x2
// 字节: 0x5e
// 操作:  ..., slot4, slot3, slot2, slot1 ->
//       ... slot2, slot1, slot4, slot3, slot2, slot1 ->
// PS: slot3 和 slot4 可以是一个操作数（double、long），也可以是两个
//
type Dup2X2 struct {
	base.NoOperandsInstruction
}

func (this *Dup2X2) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()
	slot1 := operandStack.PopSlot()
	slot2 := operandStack.PopSlot()
	slot3 := operandStack.PopSlot()
	slot4 := operandStack.PopSlot()
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
	operandStack.PushSlot(slot4)
	operandStack.PushSlot(slot3)
	operandStack.PushSlot(slot2)
	operandStack.PushSlot(slot1)
}
