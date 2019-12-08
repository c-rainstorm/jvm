package rtda

import "jvm/pkg/rtda/heap"

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
	method       *heap.Method
}

func (this *Frame) Thread() *Thread {
	return this.thread
}

func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}

func (this *Frame) LocalVars() *LocalVars {
	return &this.localVars
}

func (this *Frame) NextPC() int {
	return this.nextPC
}

func (this *Frame) SetNextPC(nextPc int) {
	this.nextPC = nextPc
}

func (this *Frame) Method() *heap.Method {
	return this.method
}
