package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

type InvokeStatic struct {
	base.Index16Instruction
}

func (this *InvokeStatic) Execute(frame *rtda.Frame) {
	methodRef := frame.Method().Class().ConstantPool().GetConstant(this.Index).(*heap.MethodSymRef)

	method := methodRef.ResolvedMethod()

	class := method.Class()
	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	if !method.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	invoke.InvokeMethod(frame, method)
}
