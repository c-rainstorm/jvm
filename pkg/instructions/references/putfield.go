package references

import (
	"jvm/pkg/global"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type PutField struct {
	base.Index16Instruction
}

func (this *PutField) Execute(frame *rtda.Frame) {
	currentMethod := frame.Method()
	currentClass := currentMethod.Class()
	fieldSymRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.FieldSymRef)
	field := fieldSymRef.ResolvedField()

	if field.IsStatic() {
		panic("java.lang.IncompatibleChangeError")
	}

	if field.IsFinal() {
		if currentClass != field.Class() || !currentMethod.IsInit() {
			panic("java.lang.IllegalAccessError")
		}
	}

	slotId := field.SlotId()
	operandStack := frame.OperandStack()

	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		val := operandStack.PopInt()
		ref := operandStack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.FieldSlots().SetInt(slotId, val)
	case global.FdFloat:
		val := operandStack.PopFloat()
		ref := operandStack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.FieldSlots().SetFloat(slotId, val)
	case global.FdLong:
		val := operandStack.PopLong()
		ref := operandStack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.FieldSlots().SetLong(slotId, val)
	case global.FdDouble:
		val := operandStack.PopDouble()
		ref := operandStack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.FieldSlots().SetDouble(slotId, val)
	case global.FdRef, global.FdArray:
		val := operandStack.PopRef()
		ref := operandStack.PopRef()
		if ref == nil {
			panic("java.lang.NullPointerException")
		}
		ref.FieldSlots().SetRef(slotId, val)
	}
}
