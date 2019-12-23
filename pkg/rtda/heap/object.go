package heap

import (
	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

type Object struct {
	class *Class
	data  interface{}
}

func (this *Object) FieldSlots() Slots {
	return this.data.(Slots)
}

func (this *Object) IsInstanceOf(targetClass *Class) bool {
	return this.isAssignableTo(targetClass)
}

func (this *Object) isAssignableTo(targetClass *Class) bool {
	if this.class == targetClass {
		return true
	}

	if this.class.isArray() {
		if targetClass.IsInterface() {
			return global.JavaIOSerializable == targetClass.name ||
				global.JavaLangCloneable == targetClass.name
		} else if targetClass.isArray() {
			thisEle := this.class.ElementClass()
			targetEle := targetClass.ElementClass()
			return thisEle == targetEle || thisEle.IsSubClassOf(targetEle)
		} else {
			return global.JavaLangObject == targetClass.name
		}
	} else if this.class.IsInterface() {
		if targetClass.IsInterface() {
			return this.class.IsSubInterfaceOf(targetClass)
		} else {
			return global.JavaLangObject == targetClass.name
		}
	} else {
		if targetClass.IsInterface() {
			return this.class.IsImplClassOf(targetClass)
		} else {
			return this.class.IsSubClassOf(targetClass)
		}
	}
}

func (this *Object) Class() *Class {
	return this.class
}

func (this *Object) ArrayLength() int32 {
	if !this.class.isArray() {
		panic("当前对象不是数组：" + this.class.name)
	}

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
	case []*Object:
		return int32(len(this.data.([]*Object)))
	default:
		panic("Not Array")
	}
}

func (this *Object) Get(index int32) interface{} {
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
	case []*Object:
		return this.data.([]*Object)[index]
	default:
		panic("Not Array")
	}
}

func (this *Object) Set(index int32, value interface{}) {
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
	case []*Object:
		this.data.([]*Object)[index] = value.(*Object)
	default:
		panic("Not Array")
	}
}

func (this *Object) indexCheck(index int32) {
	length := this.ArrayLength()
	if index < 0 || length <= index {
		panic("ArrayIndexOutOfBoundsException")
	}
}
