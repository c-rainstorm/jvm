package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
	thread       *Thread
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
