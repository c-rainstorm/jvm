package references

import (
	"jvm/pkg/global"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type GetField struct {
	base.Index16Instruction
}

func (this *GetField) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	fieldSymRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.FieldSymRef)
	field := fieldSymRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleChangeError")
	}

	slotId := field.SlotId()
	operandStack := frame.OperandStack()

	ref := operandStack.PopNormalObject()
	if ref == nil {
		panic("java.lang.NullPointerException")
	}



	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		operandStack.PushInt(ref.FieldSlots().GetInt(slotId))
	case global.FdFloat:
		operandStack.PushFloat(ref.FieldSlots().GetFloat(slotId))
	case global.FdLong:
		operandStack.PushLong(ref.FieldSlots().GetLong(slotId))
	case global.FdDouble:
		operandStack.PushDouble(ref.FieldSlots().GetDouble(slotId))
	case global.FdRef, global.FdArray:
		operandStack.PushRef(ref.FieldSlots().GetRef(slotId))
	}
}
