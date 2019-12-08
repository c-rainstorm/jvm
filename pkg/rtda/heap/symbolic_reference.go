package heap

// 符号引用
type SymbolicRef struct {
	// 进行符号解析时的那个类的运行时常量池
	cp        *ConstantPool
	className string
	class     *Class
}

func (this *SymbolicRef) ResolvedClass() *Class {
	if this.class == nil {
		this.ResolveClassRef()
	}

	return this.class
}

func (this *SymbolicRef) ResolveClassRef() *Class {
	// 使用解析符号引用的类进行类加载
	currentClass := this.cp.class
	newClass := currentClass.classLoader.LoadClass(this.className)
	if !newClass.isAccessibleTo(currentClass) {
		panic("java.lang.IllegalAccessError")
	}

	this.class = newClass
	return newClass
}
