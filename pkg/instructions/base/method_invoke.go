package base

import (
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func InvokeMethod(invokerFrame *rtda.Frame, method *heap.Method) {
	thread := invokerFrame.Thread()

	newFrame := thread.NewFrame(method)
	thread.PushFrame(newFrame)

	argSlotCount := int(method.ArgSlotCount())
	if argSlotCount > 0 {
		// 实例方法绝对大于0，因为有 this 指针
		// 静态方法可能为0
		for i := argSlotCount - 1; i >= 0; i-- {
			newFrame.LocalVars().SetSlot(uint(i), invokerFrame.OperandStack().PopSlot())
		}
	}
}
