package heap

import (
	"math"
)

type Slot struct {
	num int32
	ref Object
}

type Slots []Slot

func newSlots(slotCount uint) Slots {
	return make([]Slot, slotCount)
}

func (this Slots) SetInt(index uint, val int32) {
	this[index].num = val
}

func (this Slots) GetInt(index uint) int32 {
	return this[index].num
}

func (this Slots) SetFloat(index uint, val float32) {
	this[index].num = int32(math.Float32bits(val))
}

func (this Slots) GetFloat(index uint) float32 {
	return math.Float32frombits(uint32(this[index].num))
}

func (this Slots) SetLong(index uint, val int64) {
	this[index].num = int32(uint32(val >> 32))
	this[index+1].num = int32(uint32(val))
}

func (this Slots) GetLong(index uint) int64 {
	return (int64(uint32(this[index].num)) << 32) | int64(uint32(this[index+1].num))
}

func (this Slots) SetDouble(index uint, val float64) {
	this.SetLong(index, int64(math.Float64bits(val)))
}

func (this Slots) GetDouble(index uint) float64 {
	return math.Float64frombits(uint64(this.GetLong(index)))
}

func (this Slots) SetRef(index uint, ref Object) {
	this[index].ref = ref
}

func (this Slots) GetRef(index uint) Object {
	return this[index].ref
}
