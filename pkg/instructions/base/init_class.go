package base

import (
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func InitClass(thread *rtda.Thread, class *heap.Class) {
	if class == nil || class.InitStarted() || class.IsInterface() {
		return
	}

	class.StartInit()
	scheduleClinit(thread, class)
	InitClass(thread, class.SuperClass())
}

func scheduleClinit(thread *rtda.Thread, class *heap.Class) {
	clinit := class.GetClinitMethod()
	if clinit != nil {
		thread.PushFrame(thread.NewFrame(clinit))
	}
}
