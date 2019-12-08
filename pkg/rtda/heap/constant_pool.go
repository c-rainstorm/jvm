package heap

type Constant interface {
}

type ConstantPool struct {
	class  *Class
	consts []Constant
}

func (this *ConstantPool) GetConstant(index uint) Constant {
	if c := this.consts[index]; c != nil {
		return c
	}
	return nil
}

func (this *ConstantPool) Class() *Class {
	return this.class
}
