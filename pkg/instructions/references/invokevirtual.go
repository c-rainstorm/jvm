package references

import (
	"fmt"

	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

type InvokeVirtual struct {
	base.Index16Instruction
}

func (this *InvokeVirtual) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	methodRef := currentClass.ConstantPool().GetConstant(this.Index).(*heap.MethodSymRef)

	if methodRef.Name() == "println" {
		operandStack := frame.OperandStack()
		switch methodRef.Descriptor() {
		case "(Z)V":
			fmt.Printf("%v\n", operandStack.PopInt() != 0)
		case "(B)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(C)V":
			fmt.Printf("%c\n", operandStack.PopInt())
		case "(S)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(I)V":
			fmt.Printf("%v\n", operandStack.PopInt())
		case "(J)V":
			fmt.Printf("%v\n", operandStack.PopLong())
		case "(F)V":
			fmt.Printf("%v\n", operandStack.PopFloat())
		case "(D)V":
			fmt.Printf("%v\n", operandStack.PopDouble())
		case "(Ljava/lang/String;)V":
			fmt.Printf("%v\n", heap.GoString(operandStack.PopNormalObject()))
		default:
			panic("println: " + methodRef.Descriptor())
		}

		frame.OperandStack().PopRef()
		return
	}

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

	// 动态调用
	methodToBeInvoked := heap.LookupMethodInClass(ref.Class(), methodRef.Name(), methodRef.Descriptor())

	if methodToBeInvoked == nil || methodToBeInvoked.IsAbstract() {
		panic("java.lang.AbstractMethodError")
	}

	invoke.InvokeMethod(frame, methodToBeInvoked)
}
