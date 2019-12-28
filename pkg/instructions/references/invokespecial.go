package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

// 1. 实例的构造函数
// 2. 私有方法
// 3. super 调用父类函数
type InvokeSpecial struct {
	base.Index16Instruction
}

func (this *InvokeSpecial) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	methodRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.MethodSymRef)

	resolvedClass := methodRef.ResolvedClass()
	resolvedMethod := methodRef.ResolvedMethod()

	if resolvedMethod.IsStatic() {
		panic("java.lang.IncompatibleClassChangeError")
	}

	ref := frame.OperandStack().GetRefFromTop(resolvedMethod.ArgSlotCount())
	if ref == nil {
		panic("java.lang.NullPointerException")
	}

	// JVMS8 page483 定义的访问错误
	// 若 resolvedMethod 是当前类父类里的方法，且当前类和父类不再同一个包时，对象的类应该是当前类的子类，否则抛 IllegalAccessError 异常
	if resolvedMethod.IsProtected() &&
		currentClass.IsSubClassOf(resolvedClass) && currentClass.PackageName() != resolvedClass.PackageName() &&
		!ref.IsInstanceOf(currentClass) {
		// 方法对当前方法可见
		panic("java.lang.IllegalAccessError")
	}

	// JVMS8 page481
	// 实例方法里使用 super 调用父类方法
	methodToBeInvoked := resolvedMethod
	if currentClass.IsSuper() &&
		currentClass.IsSubClassOf(resolvedClass) &&
		!resolvedMethod.IsInit() {
		methodToBeInvoked = heap.LookupMethodInClass(currentClass.SuperClass(), methodRef.Name(), methodRef.Descriptor())
	}

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	invoke.InvokeMethod(frame, resolvedMethod)
}
