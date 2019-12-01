package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

func jumpIfSatisfyI(apply func(val int32) bool, frame *rtda.Frame, offset int) {
	val := frame.OperandStack().PopInt()
	if apply(val) {
		base.BranchJump(frame, offset)
	}
}

// ..., value ->   // 如果 value 与 0 的关系和指令相匹配，则进行分支跳转
// ...
type IfEQ struct {
	base.BranchInstruction
}

func (this *IfEQ) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val == 0
	}, frame, this.Offset)
}

type IfNE struct {
	base.BranchInstruction
}

func (this *IfNE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val != 0
	}, frame, this.Offset)
}

type IfLT struct {
	base.BranchInstruction
}

func (this *IfLT) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val < 0
	}, frame, this.Offset)
}

type IfLE struct {
	base.BranchInstruction
}

func (this *IfLE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val <= 0
	}, frame, this.Offset)
}

type IfGT struct {
	base.BranchInstruction
}

func (this *IfGT) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val > 0
	}, frame, this.Offset)
}

type IfGE struct {
	base.BranchInstruction
}

func (this *IfGE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyI(func(val int32) bool {
		return val >= 0
	}, frame, this.Offset)
}
