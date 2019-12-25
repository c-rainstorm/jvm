package heap

type ArrayObject struct {
	BaseObject
	data interface{}
}

func (this *ArrayObject) ArrayLength() int32 {
	switch this.data.(type) {
	case []int8:
		return int32(len(this.data.([]int8)))
	case []int16:
		return int32(len(this.data.([]int16)))
	case []uint16:
		return int32(len(this.data.([]uint16)))
	case []int32:
		return int32(len(this.data.([]int32)))
	case []int64:
		return int32(len(this.data.([]int64)))
	case []float32:
		return int32(len(this.data.([]float32)))
	case []float64:
		return int32(len(this.data.([]float64)))
	case []Object:
		return int32(len(this.data.([]Object)))
	default:
		panic("Not Array")
	}
}

func (this *ArrayObject) Get(index int32) interface{} {
	this.indexCheck(index)

	switch this.data.(type) {
	case []int8:
		return this.data.([]int8)[index]
	case []int16:
		return this.data.([]int16)[index]
	case []uint16:
		return this.data.([]uint16)[index]
	case []int32:
		return this.data.([]int32)[index]
	case []int64:
		return this.data.([]int64)[index]
	case []float32:
		return this.data.([]float32)[index]
	case []float64:
		return this.data.([]float64)[index]
	case []Object:
		return this.data.([]Object)[index]
	default:
		panic("Not Array")
	}
}

func (this *ArrayObject) Set(index int32, value interface{}) {
	this.indexCheck(index)
	switch this.data.(type) {
	case []int8:
		this.data.([]int8)[index] = int8(value.(int32))
	case []int16:
		this.data.([]int16)[index] = int16(value.(int32))
	case []uint16:
		this.data.([]uint16)[index] = uint16(value.(int32))
	case []int32:
		this.data.([]int32)[index] = value.(int32)
	case []int64:
		this.data.([]int64)[index] = value.(int64)
	case []float32:
		this.data.([]float32)[index] = value.(float32)
	case []float64:
		this.data.([]float64)[index] = value.(float64)
	case []Object:
		this.data.([]Object)[index] = value.(Object)
	default:
		panic("Not Array")
	}
}

func (this *ArrayObject) indexCheck(index int32) {
	length := this.ArrayLength()
	if index < 0 || length <= index {
		panic("ArrayIndexOutOfBoundsException")
	}
}

func (this *ArrayObject) CopyTo(dest *ArrayObject, srcPos int32, destPos int32, length int32) {
	switch this.data.(type) {
	case []int8:
		_src := this.data.([]int8)[srcPos : srcPos+length]
		_dest := dest.data.([]int8)[destPos : destPos+length]
		copy(_dest, _src)
	case []int16:
		_src := this.data.([]int16)[srcPos : srcPos+length]
		_dest := dest.data.([]int16)[destPos : destPos+length]
		copy(_dest, _src)
	case []uint16:
		_src := this.data.([]uint16)[srcPos : srcPos+length]
		_dest := dest.data.([]uint16)[destPos : destPos+length]
		copy(_dest, _src)
	case []int32:
		_src := this.data.([]int32)[srcPos : srcPos+length]
		_dest := dest.data.([]int32)[destPos : destPos+length]
		copy(_dest, _src)
	case []int64:
		_src := this.data.([]int64)[srcPos : srcPos+length]
		_dest := dest.data.([]int64)[destPos : destPos+length]
		copy(_dest, _src)
	case []float32:
		_src := this.data.([]float32)[srcPos : srcPos+length]
		_dest := dest.data.([]float32)[destPos : destPos+length]
		copy(_dest, _src)
	case []float64:
		_src := this.data.([]float64)[srcPos : srcPos+length]
		_dest := dest.data.([]float64)[destPos : destPos+length]
		copy(_dest, _src)
	case []Object:
		_src := this.data.([]Object)[srcPos : srcPos+length]
		_dest := dest.data.([]Object)[destPos : destPos+length]
		copy(_dest, _src)
	default:
		panic("Not Array")
	}
}
