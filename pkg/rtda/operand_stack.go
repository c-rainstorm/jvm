package rtda

import (
	"math"

	"jvm/pkg/rtda/heap"
)

type OperandStack struct {
	size  uint
	slots []Slot
}

func NewOperandStack(maxStack uint) *OperandStack {
	if maxStack > 0 {
		return &OperandStack{
			slots: make([]Slot, maxStack),
		}
	}

	return nil
}

func (this *OperandStack) PushInt(val int32) {
	this.slots[this.size].num = val
	this.size++
}

func (this *OperandStack) PopInt() int32 {
	this.size--
	return this.slots[this.size].num
}

func (this *OperandStack) PushFloat(val float32) {
	this.slots[this.size].num = int32(math.Float32bits(val))
	this.size++
}

func (this *OperandStack) PopFloat() float32 {
	this.size--
	return math.Float32frombits(uint32(this.slots[this.size].num))
}

func (this *OperandStack) PushLong(val int64) {
	this.PushInt(int32(uint32(val)))
	this.PushInt(int32(uint32(val >> 32)))
}

func (this *OperandStack) PopLong() int64 {
	return (int64(uint32(this.PopInt())) << 32) | int64(uint32(this.PopInt()))
}

func (this *OperandStack) PushDouble(val float64) {
	this.PushLong(int64(math.Float64bits(val)))
}

func (this *OperandStack) PopDouble() float64 {
	return math.Float64frombits(uint64(this.PopLong()))
}

func (this *OperandStack) PushRef(ref heap.Object) {
	this.slots[this.size].ref = ref
	this.size++
}

func (this *OperandStack) PopRef() heap.Object {
	this.size--
	ref := this.slots[this.size].ref
	this.slots[this.size].ref = nil
	return ref
}

func (this *OperandStack) PopNormalObject() *heap.NormalObject {
	popRef := this.PopRef()
	switch popRef.(type) {
	case *heap.NormalObject:
		return popRef.(*heap.NormalObject)
	case *heap.ClassObject:
		return popRef.(*heap.ClassObject).NormalObject
	default:
		panic("ref not valid")
	}
}

func (this *OperandStack) PushSlot(slot Slot) {
	this.slots[this.size] = slot
	this.size++
}

func (this *OperandStack) PopSlot() Slot {
	this.size--
	return this.slots[this.size]
}

func (this *OperandStack) GetRefFromTop(count uint) heap.Object {
	return this.slots[this.size-count].ref
}

func (this *OperandStack) Clear() {
	for _, slot := range this.slots {
		slot.num = 0
		slot.ref = nil
	}
	this.size = 0
}
