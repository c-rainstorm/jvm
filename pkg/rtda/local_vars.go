package rtda

import (
	"math"

	"jvm/pkg/rtda/heap"
)

type Slot struct {
	num int32
	ref *heap.Object
}

type LocalVars []Slot

func NewLocalVars(maxLocals uint) LocalVars {
	if maxLocals > 0 {
		return make([]Slot, maxLocals)
	}
	return nil
}

func (this LocalVars) SetInt(index uint, val int32) {
	this[index].num = val
}

func (this LocalVars) GetInt(index uint) int32 {
	return this[index].num
}

func (this LocalVars) SetFloat(index uint, val float32) {
	this[index].num = int32(math.Float32bits(val))
}

func (this LocalVars) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(this[index].num))
}

func (this LocalVars) SetLong(index uint, val int64) {
	this[index].num = int32(uint32(val))
	this[index+1].num = int32(uint32(val >> 32))
}

func (this LocalVars) GetLong(index uint) int64 {
	return int64(uint32(this[index].num)) | (int64(uint32(this[index+1].num)) << 32)
}

func (this LocalVars) SetDouble(index uint, val float64) {
	this.SetLong(index, int64(math.Float64bits(val)))
}

func (this LocalVars) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(this.GetLong(index)))
}

func (this LocalVars) SetRef(index uint, ref *heap.Object) {
	this[index].ref = ref
}

func (this LocalVars) GetRef(index uint) *heap.Object {
	return this[index].ref
}

func (this LocalVars) SetSlot(index uint, slot Slot) {
	this[index] = slot
}
