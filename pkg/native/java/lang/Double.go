package lang

import (
	"math"

	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry("java/lang/Double", "doubleToRawLongBits", "(D)J", doubleToRawLongBits)
	native.MethodRegistry.Registry("java/lang/Double", "longBitsToDouble", "(J)D", longBitsToDouble)
}

// public static native double longBitsToDouble(long bits);
func longBitsToDouble(frame *rtda.Frame) {
	longVal := frame.LocalVars().GetLong(0)
	frame.OperandStack().PushDouble(math.Float64frombits(uint64(longVal)))
}

// public static native long doubleToRawLongBits(double value);
func doubleToRawLongBits(frame *rtda.Frame) {
	floatVal := frame.LocalVars().GetDouble(0)
	frame.OperandStack().PushLong(int64(math.Float64bits(floatVal)))
}
