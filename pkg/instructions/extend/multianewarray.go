package extend

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type MultiANewArray struct {
	index      uint16
	dimensions uint8
}

func (this *MultiANewArray) FetchOperands(reader *base.ByteCodeReader) {
	this.index = reader.ReadUint16()
	this.dimensions = reader.ReadUint8()
}

func (this *MultiANewArray) Execute(frame *rtda.Frame) {
	operandStack := frame.OperandStack()

	classRef := frame.Method().Class().ConstantPool().GetConstant(uint(this.index)).(*heap.ClassSymRef)
	multiDimensionArrayClass := classRef.ResolvedClass()
	counts := popAndCheckCounts(operandStack, this.dimensions)

	operandStack.PushRef(newMultiDimensionArray(counts, multiDimensionArrayClass))
}

func newMultiDimensionArray(counts []int32, multiDimensionArrayClass *heap.ClassObject) *heap.ArrayObject {
	count := counts[0]
	array := multiDimensionArrayClass.NewArray(count)

	if len(counts) > 1 {
		for i := int32(0); i < count; i++ {
			array.Set(i, newMultiDimensionArray(counts[1:], multiDimensionArrayClass.ElementClass()))
		}
	}

	return array
}

func popAndCheckCounts(stack *rtda.OperandStack, dimensions uint8) []int32 {
	counts := make([]int32, dimensions)
	for i := int(dimensions - 1); i >= 0; i-- {
		counts[i] = stack.PopInt()
		if counts[i] < 0 {
			panic("java.lang.NegativeArraySizeException")
		}
	}
	return counts
}
