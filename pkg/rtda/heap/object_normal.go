package heap

type NormalObject struct {
	BaseObject
	slots Slots
}

func (this *NormalObject) FieldSlots() Slots {
	return this.slots
}
