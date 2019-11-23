package classfile

const (
	AttrCode               = "Code"
	AttrConstantValue      = "ConstantValue"
	AttrDeprecated         = "Deprecated"
	AttrExceptions         = "Exceptions"
	AttrLineNumberTable    = "LineNumberTable"
	AttrLocalVariableTable = "LocalVariableTable"
	AttrSourceFile         = "SourceFile"
	AttrSynthetic          = "Synthetic"
)

type AttributeInfo interface {
	read(reader *ClassReader)
}

func readAttributes(reader *ClassReader, constantPool ConstantPool) []AttributeInfo {
	attrCount := reader.readUnit16()
	attrs := make([]AttributeInfo, attrCount)

	for i := range attrs {
		attrs[i] = readAttribute(reader, constantPool)
	}

	return attrs
}

func readAttribute(reader *ClassReader, cp ConstantPool) AttributeInfo {
	attrNameIndex := reader.readUnit16()
	attrLen := reader.readUint32()

	attrName := cp[attrNameIndex].Val()
	attrInfo := newAttributeInfo(attrName.(string), attrLen, cp)

	attrInfo.read(reader)
	return attrInfo
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case AttrCode:
		return &CodeAttribute{cp: cp}
	case AttrConstantValue:
		return &ConstantValueAttribute{}
	case AttrDeprecated:
		return &DeprecatedAttribute{}
	case AttrExceptions:
		return &ExceptionAttribute{}
	case AttrLineNumberTable:
		return &LineNumberTableAttribute{}
	case AttrLocalVariableTable:
		return &LocalVariableTableAttribute{cp: cp}
	case AttrSourceFile:
		return &SourceFileAttribute{cp: cp}
	case AttrSynthetic:
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{
			attrName:  attrName,
			attrLen:   attrLen,
			attrValue: nil,
		}
	}
}

// 定长属性，只会出现在 field_info 结构里
type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (this ConstantValueAttribute) read(reader *ClassReader) {
	this.constantValueIndex = reader.readUnit16()
}

type LineNumberTableEntry struct {
	// code下标
	startPc uint16
	// 源文件行号
	lineNumber uint16
}

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}

func (this LineNumberTableAttribute) read(reader *ClassReader) {
	lineNumberTableLength := reader.readUnit16()
	this.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)

	for i := range this.lineNumberTable {
		this.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUnit16(),
			lineNumber: reader.readUnit16(),
		}
	}
}

type LocalVariableTableEntry struct {
	// code[] 下标, 必须对应一个字节码指令
	startPc uint16

	// [startPc, startPc+length) 代表变量作用域
	length uint16

	// 变量名下标，对应常量池 UTF8 类型常量
	nameIndex uint16

	// 变量描述符下标，对应常量池 UTF8 类型常量
	descriptorIndex uint16

	// 当前变量在当前栈帧变量表的索引
	index uint16
}

type LocalVariableTableAttribute struct {
	cp                 ConstantPool
	localVariableTable []*LocalVariableTableEntry
}

func (this LocalVariableTableAttribute) read(reader *ClassReader) {
	lineVariableTableLength := reader.readUnit16()
	this.localVariableTable = make([]*LocalVariableTableEntry, lineVariableTableLength)

	for i := range this.localVariableTable {
		this.localVariableTable[i] = &LocalVariableTableEntry{
			startPc:         reader.readUnit16(),
			length:          reader.readUnit16(),
			nameIndex:       reader.readUnit16(),
			descriptorIndex: reader.readUnit16(),
			index:           reader.readUnit16(),
		}
	}
}

type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (this SourceFileAttribute) read(reader *ClassReader) {
	this.sourceFileIndex = reader.readUnit16()
}

// 标记属性，不做任何的数据读取
type MarkerAttribute struct {
}

func (this MarkerAttribute) read(reader *ClassReader) {
}

type DeprecatedAttribute struct {
	MarkerAttribute
}

func (this DeprecatedAttribute) read(reader *ClassReader) {
	this.MarkerAttribute.read(reader)
}

type SyntheticAttribute struct {
	MarkerAttribute
}

func (this SyntheticAttribute) read(reader *ClassReader) {
	this.MarkerAttribute.read(reader)
}

type UnparsedAttribute struct {
	attrName  string
	attrLen   uint32
	attrValue []byte
}

func (this UnparsedAttribute) read(reader *ClassReader) {
	this.attrValue = reader.readBytes(this.attrLen)
}
