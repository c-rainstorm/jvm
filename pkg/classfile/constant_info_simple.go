package classfile

import (
	"fmt"
	"math"

	"jvm/pkg/global"
)

// --------------------- 整型常量 -------------------------

type ConstantIntegerInfo struct {
	val int32
}

func (this *ConstantIntegerInfo) String() string {
	return fmt.Sprintf("{ConstantIntegerInfo: %v}", this.val)
}

func (this *ConstantIntegerInfo) read(reader *ClassReader) {
	this.val = int32(reader.readUint32())
	if global.Verbose {
		log.Infof("parsed integer constant: %v", this.val)
	}
}

// --------------------- Long型常量 -------------------------

type ConstantLongInfo struct {
	val int64
}

func (this *ConstantLongInfo) String() string {
	return fmt.Sprintf("{ConstantLongInfo: %v}", this.val)
}

func (this *ConstantLongInfo) read(reader *ClassReader) {
	this.val = int64(reader.readUint64())
	if global.Verbose {
		log.Infof("parsed long constant: %v", this.val)
	}
}

// --------------------- float 常量 -------------------------

type ConstantFloatInfo struct {
	val float32
}

func (this *ConstantFloatInfo) String() string {
	return fmt.Sprintf("{ConstantFloatInfo: %v}", this.val)
}

func (this *ConstantFloatInfo) read(reader *ClassReader) {
	intVal := reader.readUint32()
	this.val = math.Float32frombits(intVal)
	if global.Verbose {
		log.Infof("parsed float constant: %v", this.val)
	}
}

// --------------------- double 常量 -------------------------

type ConstantDoubleInfo struct {
	val float64
}

func (this *ConstantDoubleInfo) String() string {
	return fmt.Sprintf("{ConstantDoubleInfo: %v}", this.val)
}

func (this *ConstantDoubleInfo) read(reader *ClassReader) {
	longVal := reader.readUint64()
	this.val = math.Float64frombits(longVal)
	if global.Verbose {
		log.Infof("parsed double constant: %v", this.val)
	}
}

// --------------------- UTF8 常量 -------------------------

type ConstantUtf8Info struct {
	val string
}

func (this *ConstantUtf8Info) String() string {
	return fmt.Sprintf("{ConstantUtf8Info: %v}", this.val)
}

func (this *ConstantUtf8Info) read(reader *ClassReader) {
	length := reader.readUnit16()
	this.val = string(reader.readBytes(length))
	if global.Verbose {
		log.Infof("parsed utf8 constant: %v", this.val)
	}
}

// --------------------- String 常量 -------------------------

type ConstantStringInfo struct {
	cp    ConstantPool
	index uint16
}

func (this *ConstantStringInfo) String() string {
	return fmt.Sprintf("{ConstantStringInfo: %v}", this.cp[this.index].(*ConstantUtf8Info).val)
}

func (this *ConstantStringInfo) read(reader *ClassReader) {
	this.index = reader.readUnit16()
}

// --------------------- class 常量 -------------------------

type ConstantClassInfo struct {
	cp    ConstantPool
	index uint16
}

func (this *ConstantClassInfo) String() string {
	return fmt.Sprintf("{ConstantClassInfo: %v}", this.cp[this.index].(*ConstantUtf8Info).val)
}

func (this *ConstantClassInfo) read(reader *ClassReader) {
	this.index = reader.readUnit16()
}
