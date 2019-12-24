package heap

import "jvm/pkg/global"

type NormalObject struct {
	BaseObject
	slots Slots
}

func (this *NormalObject) FieldSlots() Slots {
	return this.slots
}

func (this *NormalObject) SetField(name string, descriptor string, value interface{}) {
	field := this.Class().lookupField(name, descriptor)

	if field == nil {
		panic("field not found: " + name + " " + descriptor)
	}

	slotId := field.slotId
	var slots Slots
	if field.IsStatic() {
		slots = field.class.StaticSlots()
	} else {
		slots = this.FieldSlots()
	}

	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		slots.SetInt(slotId, value.(int32))
	case global.FdFloat:
		slots.SetFloat(slotId, value.(float32))
	case global.FdLong:
		slots.SetLong(slotId, value.(int64))
	case global.FdDouble:
		slots.SetDouble(slotId, value.(float64))
	case global.FdRef, global.FdArray:
		slots.SetRef(slotId, value.(Object))
	}
}

func (this *NormalObject) GetField(name string, descriptor string) interface{} {
	field := this.Class().lookupField(name, descriptor)
	if field == nil {
		panic("field not found: " + name + " " + descriptor)
	}

	slotId := field.slotId
	var slots Slots
	if field.IsStatic() {
		slots = field.class.StaticSlots()
	} else {
		slots = this.FieldSlots()
	}

	switch string(field.Descriptor()[0]) {
	case global.FdBoolean, global.FdByte, global.FdChar, global.FdShort, global.FdInt:
		return slots.GetInt(slotId)
	case global.FdFloat:
		return slots.GetFloat(slotId)
	case global.FdLong:
		return slots.GetLong(slotId)
	case global.FdDouble:
		return slots.GetDouble(slotId)
	case global.FdRef, global.FdArray:
		return slots.GetRef(slotId)
	default:
		panic("unknow type" + string(field.Descriptor()[0]))
	}
}
