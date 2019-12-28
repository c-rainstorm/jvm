package reflect

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	// public static native Class<?> getCallerClass()
	native.MethodRegistry.Registry("sun/reflect/Reflection", "getCallerClass", "()Ljava/lang/Class;", getCallerClass)
}

func getCallerClass(frame *rtda.Frame) {
	getCallerClassFrame := frame.Thread().TopFrame() // native getCallerClass frame
	frame.OperandStack().PushRef(getCallerClassFrame.Lower().Method().Class())
}
