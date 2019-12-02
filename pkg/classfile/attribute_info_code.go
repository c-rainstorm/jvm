package classfile

// 变长属性，只存在于 Method_info 里
// 如果方法是 native/abstract 对应的 MethodInfo 里不应该有该属性
// 否则，MethodInfo 里有且只能有一个 Code 属性
type CodeAttribute struct {
	cp ConstantPool
	// 操作数栈的最大深度
	maxStack uint16
	// 局部变量表大小，对于 long/double 类型的变量，最大变量下标应该是  maxLocal-2，其他变量应为 maxLocal-1
	maxLocals uint16
	// 字节码
	code []byte
	// 异常表，每个实体代表一个异常处理器，异常表里实体的顺序非常重要，影响异常捕获的逻辑
	exceptionTable []*ExceptionTableEntry
	// 属性表
	attributes []AttributeInfo
}

type ExceptionTableEntry struct {
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
	catchType uint16
}

func (this CodeAttribute) read(reader *ClassReader) {
	this.maxStack = reader.readUnit16()
	this.maxLocals = reader.readUnit16()
	this.code = reader.readBytes(reader.readUint32())
	this.exceptionTable = readExceptionTable(reader)
	this.attributes = readAttributes(reader, this.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUnit16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)

	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUnit16(),
			endPc:     reader.readUnit16(),
			handlerPc: reader.readUnit16(),
			catchType: reader.readUnit16(),
		}
	}

	return exceptionTable
}

type ExceptionAttribute struct {
	exceptionIndexTable []uint16
}

func (this ExceptionAttribute) read(reader *ClassReader) {
	this.exceptionIndexTable = reader.readUint16s(reader.readUnit16())
}

func (this *CodeAttribute) MaxStack() uint16 {
	return this.maxStack
}

func (this *CodeAttribute) MaxLocals() uint16 {
	return this.maxLocals
}

func (this *CodeAttribute) Code() []byte {
	return this.code
}
