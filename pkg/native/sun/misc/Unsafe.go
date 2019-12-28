package misc

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	// public native int arrayBaseOffset(Class<?> var1)
	native.MethodRegistry.Registry("sun/misc/Unsafe", "arrayBaseOffset", "(Ljava/lang/Class;)I", arrayBaseOffset)
	//    public native int arrayIndexScale(Class<?> var1);
	native.MethodRegistry.Registry("sun/misc/Unsafe", "arrayIndexScale", "(Ljava/lang/Class;)I", arrayIndexScale)
	// public native int addressSize();
	native.MethodRegistry.Registry("sun/misc/Unsafe", "addressSize", "()I", addressSize)
}

func addressSize(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func arrayIndexScale(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

func arrayBaseOffset(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}
