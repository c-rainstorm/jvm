package base

import "jvm/pkg/rtda"

type Instruction interface {
	// 从 reader 里读取操作数，不同的指令操作数个数、类型不同
	FetchOperands(reader *ByteCodeReader)

	// 执行指令
	Execute(frame *rtda.Frame)
}

// 无操作数的指令
type NoOperandsInstruction struct {
}

func (this *NoOperandsInstruction) FetchOperands(reader *ByteCodeReader) {
	// 本身没有操作数，do nothing
}

// 操作数是单字节下标，通常是从本地变量表的下标
type Index8Instruction struct {
	Index uint8
}

func (this *Index8Instruction) FetchOperands(reader *ByteCodeReader) {
	this.Index = reader.ReadUint8()
}
