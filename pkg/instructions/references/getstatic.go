package references

import (
	"jvm/pkg/global"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type GetStatic struct {
	base.Index16Instruction
}

func (this *GetStatic) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	fieldSymRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.FieldSymRef)
	field := fieldSymRef.ResolvedField()

	if !field.IsStatic() {
		panic("java.lang.IncompatibleChangeError")
	}

	slotId := field.SlotId()
	staticSlots := field.Class().StaticSlots()
	operandStack := frame.OperandStack()

	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		operandStack.PushInt(staticSlots.GetInt(slotId))
	case global.FdFloat:
		operandStack.PushFloat(staticSlots.GetFloat(slotId))
	case global.FdLong:
		operandStack.PushLong(staticSlots.GetLong(slotId))
	case global.FdDouble:
		operandStack.PushDouble(staticSlots.GetDouble(slotId))
	case global.FdRef, global.FdArray:
		operandStack.PushRef(staticSlots.GetRef(slotId))
	}
}
