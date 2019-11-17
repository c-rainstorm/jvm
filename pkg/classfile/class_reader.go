package classfile

import "encoding/binary"

type ClassReader struct {
	data  []byte
	index uint16
}

func NewClassReader(data []byte) *ClassReader {
	return &ClassReader{
		data:  data,
		index: 0,
	}
}

func (this *ClassReader) readUint8() uint8 {
	val := this.data[this.index]
	this.index++
	return val
}

func (this *ClassReader) readUnit16() uint16 {
	byteSize := uint16(2)
	val := binary.BigEndian.Uint16(this.data[this.index : this.index+byteSize])
	this.index += byteSize
	return val
}

func (this *ClassReader) readUint32() uint32 {
	byteSize := uint16(4)
	val := binary.BigEndian.Uint32(this.data[this.index : this.index+byteSize])
	this.index += byteSize
	return val
}

func (this *ClassReader) readUint64() uint64 {
	byteSize := uint16(8)
	val := binary.BigEndian.Uint64(this.data[this.index : this.index+byteSize])
	this.index += byteSize
	return val
}

func (this *ClassReader) readUint16s(num uint16) []uint16 {
	s := make([]uint16, num)
	for i := range s {
		s[i] = this.readUnit16()
	}
	return s
}

func (this *ClassReader) readBytes(num uint16) []byte {
	val := this.data[this.index : this.index+num]
	this.index += num
	return val
}
