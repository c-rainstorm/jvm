package lang

import (
	"jvm/pkg/native"
	"jvm/pkg/native/sun/misc"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

func init() {
	native.MethodRegistry.Registry("java/lang/System", "arraycopy", "(Ljava/lang/Object;ILjava/lang/Object;II)V", arrayCopy)
	native.MethodRegistry.Registry("java/lang/System", "initProperties", "(Ljava/util/Properties;)Ljava/util/Properties;", initProperties)
	// private static native void setOut0(PrintStream out);
	native.MethodRegistry.Registry("java/lang/System", "setOut0", "(Ljava/io/PrintStream;)V", setOut0)

}

func setOut0(frame *rtda.Frame) {
	SystemClass := frame.Method().Class().ClassLoader().LoadClass("java/lang/System")
	outStream := frame.LocalVars().GetRef(0).(*heap.NormalObject)
	SystemClass.SetField(nil, "out", "Ljava/io/PrintStream;", outStream)
}

func initProperties(frame *rtda.Frame) {
	props := frame.LocalVars().GetRef(0).(*heap.NormalObject)
	frame.OperandStack().PushRef(props)

	for key, value := range misc.Properties {
		invoke.SetProperty(frame.Thread(), props, key, value)
	}
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
