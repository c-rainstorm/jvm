package base

type ByteCodeReader struct {
	code []byte
	pc   int
}

func (this *ByteCodeReader) Reset(code []byte, pc int) {
	this.code = code
	this.pc = pc
}

// 从当前PC指向的的位置读取一个字节
func (this *ByteCodeReader) ReadUint8() uint8 {
	byteData := this.code[this.pc]
	this.pc++
	return byteData
}

// 字节码字节流大段法存储，先读到的数据在高位
func (this *ByteCodeReader) ReadUint16() uint16 {
	return (uint16(this.ReadUint8()) << 8) | uint16(this.ReadUint8())
}

func (this *ByteCodeReader) ReadInt16() int16 {
	return int16(this.ReadUint16())
}

// 字节码字节流大段法存储，先读到的数据在高位
func (this *ByteCodeReader) ReadInt32() int32 {
	byte1 := int32(this.ReadUint8())
	byte2 := int32(this.ReadUint8())
	byte3 := int32(this.ReadUint8())
	byte4 := int32(this.ReadUint8())
	return (byte1 << 24) | (byte2 << 16) | (byte3 << 8) | byte4
}
