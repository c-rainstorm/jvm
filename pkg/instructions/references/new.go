package references

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
//
// 格式: new [indexbyte1] [indexbyte2]
// 字节: 0xbb 0x00 0x01
//   指令样例     8进制表示          指令含义
//   new 1       0xbb 0x00 0x01    使用当前类运行时常量池下标为 i 的地方的类符号引用，实例化一个新的对象，并将引用推送到操作数的栈顶
// 操作: 使用当前类运行时常量池下标为 <index> 的地方的类符号引用，实例化一个新的对象，并将引用推送到操作数的栈顶
//
type New struct {
	base.Index16Instruction
}

func (this *New) Execute(frame *rtda.Frame) {
	currentClass := frame.Method().Class()
	cp := currentClass.ConstantPool()
	// 当前帧的方法可能创建的是其他类的对象
	// 所以这块只能进行 resolve 一下
	classRef := cp.GetConstant(this.Index).(*heap.ClassSymRef)
	class := classRef.ResolvedClass()

	if class.IsInterface() || class.IsAbstract() {
		panic("java.lang.InstantiationError")
	}

	if !class.InitStarted() {
		frame.RevertNextPC()
		base.InitClass(frame.Thread(), class)
		return
	}

	objRef := class.NewObject()
	frame.OperandStack().PushRef(objRef)
}
