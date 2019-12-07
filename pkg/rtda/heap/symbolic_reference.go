package heap

// 符号引用
type SymbolicRef struct {
	cp        *ConstantPool
	className string
	class     *Class
}
