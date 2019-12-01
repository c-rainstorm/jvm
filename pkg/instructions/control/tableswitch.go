package control

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 操作数在 [low, high] 之间
type TableSwitch struct {
	// 默认的跳转偏移量
	defaultOffset int32
	// 最小的匹配操作数
	low int32
	// 最大的匹配操作数
	high int32
	// 每个操作数对应的跳转偏移量
	jumpOffsets []int32
}

func (this *TableSwitch) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.low = reader.ReadInt32()
	this.high = reader.ReadInt32()
	this.jumpOffsets = make([]int32, this.high-this.low+1)
	for i := range this.jumpOffsets {
		this.jumpOffsets[i] = reader.ReadInt32()
	}
}

func (this *TableSwitch) Execute(frame *rtda.Frame) {
	index := frame.OperandStack().PopInt()

	if index >= this.low && index <= this.high {
		base.BranchJump(frame, int(this.jumpOffsets[index-this.low]))
	} else {
		base.BranchJump(frame, int(this.defaultOffset))
	}
}
