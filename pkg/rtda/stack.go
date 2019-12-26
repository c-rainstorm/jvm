package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func (this *Stack) push(frame *Frame) {
	if this.size >= this.maxSize {
		log.Panic("java.lang.StackOverflowError")
	}

	if this._top == nil {
		frame.lower = nil
	} else {
		frame.lower = this._top
	}

	this._top = frame
	this.size++
}

func (this *Stack) pop() *Frame {
	if this._top == nil {
		log.Panic("jvm stack is empty!")
	}

	top := this._top
	this._top = top.lower
	top.lower = nil
	this.size--

	return top
}

func (this *Stack) top() *Frame {
	if this._top == nil {
		log.Panic("jvm stack is empty!")
	}

	return this._top
}

func NewStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
