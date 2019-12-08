package heap

type Method struct {
	ClassMember
	// 所需栈容量
	maxStack uint
	// 方法局部变量表大小
	maxLocals uint
	// 方法字节码
	code []byte
}

func (this *Method) MaxLocals() uint {
	return this.maxLocals
}

func (this *Method) MaxStack() uint {
	return this.maxLocals
}

func (this *Method) Code() []byte {
	return this.code
}

func (this *Method) Class() *Class {
	return this.class
}

func (this *Method) IsCLInit() bool {
	return this.hasFlag(ACC_STATIC) && this.name == "<clinit>"
}

func (this *Method) IsInit() bool {
	return !this.hasFlag(ACC_STATIC) && this.name == "<init>"
}
