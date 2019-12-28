package heap

import (
	"jvm/pkg/classfile"
	"jvm/pkg/global"
)

// ExceptionTableEntry
type ExceptionHandler struct {
	// startPc 和 endPc 代表出异常时执行的代码范围
	// startPc 是 code[] 的起始下标，endPc 是 code[] 的结束下标
	// [startPc, endPc)   实际的范围不包括 endPc 下标锁代表的位置
	// 不包括 endPC 是 JVM规范设计上的一个缺陷，具体为啥没搞明白，，，
	startPc uint16
	endPc   uint16

	// code[] 里异常处理器的下标位置，且必须指向一个字节码指令
	handlerPc uint16

	// 非0，必须是常量池的有效索引，且必须是一个 Class 类型的常量信息
	// 如果为0，则代表对所有异常都适用，通常用来实现 finally
	// finally 时该字段为 nil
	catchType *ClassSymRef
}

func (this *ExceptionHandler) StartPc() uint16 {
	return this.startPc
}

func (this *ExceptionHandler) EndPc() uint16 {
	return this.endPc
}

func (this *ExceptionHandler) HandlerPc() uint16 {
	return this.handlerPc
}

type ExceptionTable []*ExceptionHandler

type Method struct {
	ClassMember
	// 所需栈容量
	maxStack uint
	// 方法局部变量表大小
	maxLocals uint
	// 方法字节码
	code []byte
	// 方法参数占用槽数
	argSlotCount uint
	// 异常表
	exceptionTable ExceptionTable
	// 行号表
	lineNumberTable []*classfile.LineNumberTableEntry
}

func (this *Method) MaxLocals() uint {
	return this.maxLocals
}

func (this *Method) MaxStack() uint {
	return this.maxStack
}

func (this *Method) Code() []byte {
	return this.code
}

func (this *Method) Class() *ClassObject {
	return this.class
}

func (this *Method) IsCLInit() bool {
	return this.hasFlag(ACC_STATIC) && this.name == "<clinit>"
}

func (this *Method) IsInit() bool {
	return !this.hasFlag(ACC_STATIC) && this.name == "<init>"
}

func (this *Method) ArgSlotCount() uint {
	return this.argSlotCount
}

// 方法描述符格式 (方法参数列表)返回描述符
// 样例 (D[D[[D[Ljava.lang.String;Ljava.lang.String;D)V  -> void (double, double[]. double[][], String[], double)
func (this *Method) calArgSlotCount(parsedDescriptor MethodDescriptor) {
	this.argSlotCount = parsedDescriptor.getParamSlotCount()
	// 实例方法多一个 this
	// 通过接口调用的
	if !this.IsStatic() {
		this.argSlotCount++
	}

	if global.Verbose {
		log.Infof("arg slot count of %s.%s%s is %d", this.class.name, this.name, this.descriptor, this.argSlotCount)
	}
}

func (this *Method) IsAbstract() bool {
	return this.hasFlag(ACC_ABSTRACT)
}

func (this *Method) IsNative() bool {
	return this.hasFlag(ACC_NATIVE)
}

func VoidVirtualMethod(maxStack, maxLocals uint) *Method {
	return &Method{
		ClassMember: ClassMember{name: "VoidVirtualMethod", descriptor: "()V", class: &ClassObject{name: "VoidVirtualClass"}},
		maxStack:    maxStack,
		maxLocals:   maxLocals,
		code:        []byte{global.OpcReturn},
	}
}

func (this *Method) InjectNativeCodeAttr(returnType string) {
	this.maxStack = 4
	this.maxLocals = this.ArgSlotCount()
	switch string(returnType[0]) {
	case "V":
		this.code = []byte{0xFE, 0xB1} // return
	case global.FdDouble:
		this.code = []byte{0xFE, 0xAF} // dreturn
	case global.FdFloat:
		this.code = []byte{0xFE, 0xAE} // freturn
	case global.FdLong:
		this.code = []byte{0xFE, 0xAD} // lreturn
	case global.FdRef:
		this.code = []byte{0xFE, 0xB0} // areturn
	default:
		this.code = []byte{0xFE, 0xAC} // ireturn
	}
}

func (this *Method) LookupExceptionTable(class *ClassObject, pc uint16) *ExceptionHandler {
	for _, handler := range this.exceptionTable {
		pcValid := pc >= handler.startPc && pc < handler.endPc
		if !pcValid {
			continue
		}

		isFinallyBlock := handler.catchType == nil
		if isFinallyBlock {
			return handler
		}

		catchExceptionClass := handler.catchType.ResolvedClass()
		catchTypeMatched := class == catchExceptionClass ||
			(catchExceptionClass.IsInterface() && class.IsImplClassOf(catchExceptionClass)) ||
			(class.IsSubClassOf(catchExceptionClass))
		if catchTypeMatched {
			return handler
		}
	}
	return nil
}

func (this *Method) GetLineNumber(pc int) int {
	if this.lineNumberTable == nil {
		return -1
	}
	for i := len(this.lineNumberTable) - 1; i >= 0; i-- {
		entry := this.lineNumberTable[i]
		if pc >= int(entry.GetStartPc()) {
			return int(entry.GetLineNumber())
		}
	}
	return -1
}
