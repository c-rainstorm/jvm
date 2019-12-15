package references

import (
	"jvm/pkg/global"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

// 将栈顶元素推到当前类运行时常量池的 index 所表示的字段里, index 位置是 字段的符号引用
type PutStatic struct {
	base.Index16Instruction
}

func (this *PutStatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	fieldSymRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.FieldSymRef)
	field := fieldSymRef.ResolvedField()

	class := field.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !field.IsStatic() {
		panic("java.lang.IncompatibleChangeError")
	}

	if field.IsFinal() {
		if currentClass != field.Class() || !currentMethod.IsCLInit() {
			panic("java.lang.IllegalAccessError")
		}
	}

	slotId := field.SlotId()
	staticSlots := field.Class().StaticSlots()
	operandStack := frame.OperandStack()

	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		staticSlots.SetInt(slotId, operandStack.PopInt())
	case global.FdFloat:
		staticSlots.SetFloat(slotId, operandStack.PopFloat())
	case global.FdLong:
		staticSlots.SetLong(slotId, operandStack.PopLong())
	case global.FdDouble:
		staticSlots.SetDouble(slotId, operandStack.PopDouble())
	case global.FdRef, global.FdArray:
		staticSlots.SetRef(slotId, operandStack.PopRef())
	}
}
