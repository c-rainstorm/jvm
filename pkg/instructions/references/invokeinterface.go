package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

type InvokeInterface struct {
	index uint
}

func (this *InvokeInterface) FetchOperands(reader *base.ByteCodeReader) {
	this.index = uint(reader.ReadUint16())
	reader.ReadUint8()
	reader.ReadUint8()
}

func (this *InvokeInterface) Execute(frame *rtda.Frame) {
	methodRef := frame.Method().Class().ConstantPool().GetConstant(this.index).(*heap.InterfaceMethodSymRef)

	resolvedInterfaceMethod := methodRef.ResolvedInterfaceMethod()

	if resolvedInterfaceMethod.IsStatic() || resolvedInterfaceMethod.IsPrivate() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedInterfaceMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// 实际对象的类未实现当前接口
	if !ref.Class().IsImplClassOf(resolvedInterfaceMethod.Class()) {
		panic("java.lang.IncompatibleClassChangeError")
	}

	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	if !methodToBeInvoked.IsPublic() {
		panic("java.lang.illegalAccessError")
	}

	invoke.InvokeMethod(frame, methodToBeInvoked)
}
