package control

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type LookupSwitch struct {
	defaultOffset int32
	npairs        int32
	matchOffsets  []int32
}

func (this *LookupSwitch) FetchOperands(reader *base.ByteCodeReader) {
	reader.SkipPadding()
	this.defaultOffset = reader.ReadInt32()
	this.npairs = reader.ReadInt32()
	this.matchOffsets = make([]int32, this.npairs*2)
	for i := range this.matchOffsets {
		this.matchOffsets[i] = reader.ReadInt32()
	}
}

func (this *LookupSwitch) Execute(frame *rtda.Frame) {
	key := frame.OperandStack().PopInt()
	for i := int32(0); i < this.npairs*2; i += 2 {
		if this.matchOffsets[i] == key {
			base.BranchJump(frame, int(this.matchOffsets[i+1]))
			break
		}
	}

	base.BranchJump(frame, int(this.defaultOffset))
}
