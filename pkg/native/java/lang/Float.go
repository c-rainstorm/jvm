package lang

import (
	"math"

	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry("java/lang/Float", "floatToRawIntBits", "(F)I", floatToRawIntBits)
}

// public static native int floatToRawIntBits(float value);
func floatToRawIntBits(frame *rtda.Frame) {
	floatVal := frame.LocalVars().GetFloat(0)
	frame.OperandStack().PushInt(int32(math.Float32bits(floatVal)))
}
