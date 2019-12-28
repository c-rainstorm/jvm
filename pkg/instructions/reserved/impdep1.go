package reserved

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/native"
	"jvm/pkg/rtda"

	_ "jvm/pkg/native/java/io"
	_ "jvm/pkg/native/java/lang"
	_ "jvm/pkg/native/sun/misc"
	_ "jvm/pkg/native/sun/reflect"
)

// 使用该指令实现本地方法调用
type ImpDep1 struct {
	base.NoOperandsInstruction
}

func (this *ImpDep1) Execute(frame *rtda.Frame) {
	method := frame.Method()
	nativeMethodImpl := native.MethodRegistry.Find(method.Class().Name(), method.Name(), method.Descriptor())
	if nativeMethodImpl == nil {
		panic("java.lang.UnsatisfiedLinkError: " + method.Class().Name() + "." + method.Name() + method.Descriptor())
	}

	nativeMethodImpl(frame)
}
