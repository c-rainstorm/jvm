package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
	nextPC       int
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
