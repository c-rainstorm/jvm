package control

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 无条件进行跳转
type GOTO struct {
	base.BranchInstruction
}

func (this *GOTO) Execute(frame *rtda.Frame) {
	base.BranchJump(frame, this.Offset)
}
