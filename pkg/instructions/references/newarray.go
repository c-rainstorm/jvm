package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 创建基本类型的数组
type NewArray struct {
	aType uint8
}

func (this *NewArray) FetchOperands(reader *base.ByteCodeReader) {
	this.aType = reader.ReadUint8()
}

func (this *NewArray) Execute(frame *rtda.Frame) {
	loader := frame.Method().Class().ClassLoader()
	operandStack := frame.OperandStack()

	count := operandStack.PopInt()
	if count < 0 {
		panic("java.lang.NegativeArraySizeException")
	}

	arrayClass := loader.LoadPrimitiveArrayClass(this.aType)
	operandStack.PushRef(arrayClass.NewArray(count))
}
