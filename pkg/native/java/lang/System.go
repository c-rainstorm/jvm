package lang

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func init() {
	native.MethodRegistry.Registry("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arrayCopy)
}

func arrayCopy(frame *rtda.Frame) {
	vars := frame.LocalVars()
	src := vars.GetRef(0)
	srcPos := vars.GetInt(1)
	dest := vars.GetRef(2)
	destPos := vars.GetInt(3)
	length := vars.GetInt(4)

	if src == nil || dest == nil {
		panic("java.lang.NullPointerException")
	}

	srcArray := src.(*heap.ArrayObject)
	destArray := dest.(*heap.ArrayObject)

	if srcPos < 0 || destPos < 0 || length < 0 ||
		srcPos+length > srcArray.ArrayLength() ||
		destPos+length > destArray.ArrayLength() {
		panic("java.lang.IndexOutOfBoundsException")
	}

	if src.Class().ElementClass().IsPrimitive() ||
		dest.Class().ElementClass().IsPrimitive() {
		if src.Class() != src.Class() {
			panic("java.lang.ArrayStoreException")
		}
	}

	srcArray.CopyTo(destArray, srcPos, destPos, length)
}
