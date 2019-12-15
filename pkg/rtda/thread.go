package rtda

import (
	"jvm/pkg/logger"
	"jvm/pkg/rtda/heap"
)

var log = logger.NewLogrusLogger()

type Thread struct {
	// 程序计数器
	pc int
	// 栈帧
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: NewStack(1024),
	}
}

func (this *Thread) NewFrame(method *heap.Method) *Frame {
	return &Frame{
		localVars:    NewLocalVars(method.MaxLocals()),
		operandStack: NewOperandStack(method.MaxStack()),
		method:       method,
		thread:       this,
	}
}

func (this *Thread) PC() int {
	return this.pc
}

func (this *Thread) SetPC(newPc int) {
	this.pc = newPc
}

func (this *Thread) PushFrame(frame *Frame) {
	this.stack.push(frame)
}

func (this *Thread) PopFrame() *Frame {
	return this.stack.pop()
}

func (this *Thread) TopFrame() *Frame {
	return this.stack.top()
}

func (this *Thread) CurrentFrame() *Frame {
	return this.stack.top()
}

func (this *Thread) IsStackEmpty() bool {
	return this.stack.size == 0
}
