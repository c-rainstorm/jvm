package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

func jumpIfSatisfyII(apply func(val1, val2 int32) bool, frame *rtda.Frame, offset int) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopInt()
	val1 := operandStack.PopInt()
	if apply(val1, val2) {
		base.BranchJump(frame, offset)
	}
}

// ..., value1, value2 ->
// ...                        // 弹出两个操作数，如果和指令期待的比较结果相同则进行跳转
type IfICmpEQ struct {
	base.BranchInstruction
}

func (this *IfICmpEQ) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 == val2
	}, frame, this.Offset)
}

type IfICmpNE struct {
	base.BranchInstruction
}

func (this *IfICmpNE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 != val2
	}, frame, this.Offset)
}

type IfICmpLT struct {
	base.BranchInstruction
}

func (this *IfICmpLT) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 < val2
	}, frame, this.Offset)
}

type IfICmpLE struct {
	base.BranchInstruction
}

func (this *IfICmpLE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 <= val2
	}, frame, this.Offset)
}

type IfICmpGT struct {
	base.BranchInstruction
}

func (this *IfICmpGT) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 > val2
	}, frame, this.Offset)
}

type IfICmpGE struct {
	base.BranchInstruction
}

func (this *IfICmpGE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyII(func(val1, val2 int32) bool {
		return val1 >= val2
	}, frame, this.Offset)
}
