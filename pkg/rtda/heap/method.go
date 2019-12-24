package heap

import (
	"jvm/pkg/global"
)

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
