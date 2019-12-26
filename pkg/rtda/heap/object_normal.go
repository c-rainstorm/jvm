package heap

type NormalObject struct {
	BaseObject
	slots  Slots
	extend interface{}
}

func (this *NormalObject) FieldSlots() Slots {
	return this.slots
}

func (this *NormalObject) SetExtend(extend interface{}) {
	this.extend = extend
}

func (this *NormalObject) GetExtend() interface{} {
	return this.extend
}
