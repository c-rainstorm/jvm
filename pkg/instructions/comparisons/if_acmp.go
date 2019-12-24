package comparisons

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func jumpIfSatisfyIII(apply func(val1, val2 heap.Object) bool, frame *rtda.Frame, offset int) {
	operandStack := frame.OperandStack()
	val2 := operandStack.PopRef()
	val1 := operandStack.PopRef()
	if apply(val1, val2) {
		base.BranchJump(frame, offset)
	}
}

// ..., value2, value1 ->
// ...                        // 引用比较，如果和指令期待结果相同，则进行跳转
type IfACmpEQ struct {
	base.BranchInstruction
}

func (this *IfACmpEQ) Execute(frame *rtda.Frame) {
	jumpIfSatisfyIII(func(val1, val2 heap.Object) bool {
		return val1 == val2
	}, frame, this.Offset)
}

type IfACmpNE struct {
	base.BranchInstruction
}

func (this *IfACmpNE) Execute(frame *rtda.Frame) {
	jumpIfSatisfyIII(func(val1, val2 heap.Object) bool {
		return val1 != val2
	}, frame, this.Offset)
}
