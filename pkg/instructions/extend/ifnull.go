package extend

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

func jumpIfSatisfy(apply func(ref *rtda.Object) bool, frame *rtda.Frame, offset int) {
	operandStack := frame.OperandStack()
	ref := operandStack.PopRef()
	if apply(ref) {
		base.BranchJump(frame, offset)
	}
}

type IfNull struct {
	base.BranchInstruction
}

func (this *IfNull) Execute(frame *rtda.Frame) {
	jumpIfSatisfy(func(ref *rtda.Object) bool {
		return ref == nil
	}, frame, this.Offset)
}

type IfNonNull struct {
	base.BranchInstruction
}

func (this *IfNonNull) Execute(frame *rtda.Frame) {
	jumpIfSatisfy(func(ref *rtda.Object) bool {
		return ref != nil
	}, frame, this.Offset)
}
