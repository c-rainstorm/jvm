package lang

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func init() {
	native.MethodRegistry.Registry("java/lang/String", "intern", "()V", intern)
}

// public native String intern();
func intern(frame *rtda.Frame) {
	jStr := frame.LocalVars().GetThis()

	frame.OperandStack().PushRef(heap.Intern0(jStr.(*heap.NormalObject)))
}
