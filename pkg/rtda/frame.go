package rtda

type Frame struct {
	lower        *Frame
	localVars    LocalVars
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
	}
}

func (this *Frame) OperandStack() *OperandStack {
	return this.operandStack
}

func (this *Frame) LocalVars() *LocalVars {
	return &this.localVars
}
