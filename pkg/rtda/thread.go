package rtda

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

func (this *Thread) NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    NewLocalVars(maxLocals),
		operandStack: NewOperandStack(maxStack),
		thread: this,
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

func (this *Thread) currentFrame() *Frame {
	return this.stack.top()
}
