package extend

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type GOTOW struct {
	offset int
}

func (this *GOTOW) FetchOperands(reader *base.ByteCodeReader) {
	this.offset = int(reader.ReadInt32())
}

func (this *GOTOW) Execute(frame *rtda.Frame) {
	base.BranchJump(frame, this.offset)
}
