package classfile

import "fmt"

// --------------------- 字段引用、类方法引用、接口方法引用的公共结构 -------------------------

type ConstantMemberRefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (this *ConstantMemberRefInfo) String() string {
	return fmt.Sprintf("{ConstantMemberRefInfo: {class: %v, nameAndType: %v}}",
		this.cp[this.classIndex], this.cp[this.nameAndTypeIndex])
}

func (this *ConstantMemberRefInfo) read(reader *ClassReader) {
	this.classIndex = reader.readUnit16()
	this.nameAndTypeIndex = reader.readUnit16()
}

// --------------------- 特定域的描述符信息(字段、类方法签名、接口方法签名) -------------------------

// byte - B
// short - S
// char - C
// int - I
// long - J       **比较特殊，不是L，用来表示引用**
// float - F
// double - D
// java.lang.Object - Ljava/lang/Object;
// int[] - [I
// int[][] - [[I
// String[] - [Ljava/lang/String;
// void main(String[] args) - ([Ljava/lang/String;)V

type ConstantNameAndTypeRefInfo struct {
	cp              ConstantPool
	nameIndex       uint16
	descriptorIndex uint16
}

func (this *ConstantNameAndTypeRefInfo) Val() interface{} {
	panic("implement me")
}

func (this *ConstantNameAndTypeRefInfo) String() string {
	return fmt.Sprintf("{ConstantNameAndTypeRefInfo: {name: %v, descriptor: %v}}",
		this.cp[this.nameIndex], this.cp[this.descriptorIndex])
}

func (this *ConstantNameAndTypeRefInfo) read(reader *ClassReader) {
	this.nameIndex = reader.readUnit16()
	this.descriptorIndex = reader.readUnit16()
}

// --------------------- 字段引用 -------------------------

type ConstantFieldRefInfo struct {
	ConstantMemberRefInfo
}

func (this *ConstantFieldRefInfo) Val() interface{} {
	panic("implement me")
}

func (this *ConstantFieldRefInfo) String() string {
	return fmt.Sprintf("{ConstantFieldRefInfo: %v}", this.ConstantMemberRefInfo.String())
}

func (this *ConstantFieldRefInfo) read(reader *ClassReader) {
	this.ConstantMemberRefInfo.read(reader)
}

// --------------------- 方法引用 -------------------------

type ConstantMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (this *ConstantMethodRefInfo) Val() interface{} {
	panic("implement me")
}

func (this *ConstantMethodRefInfo) String() string {
	return fmt.Sprintf("{ConstantMethodRefInfo: %v}", this.ConstantMemberRefInfo.String())
}

func (this *ConstantMethodRefInfo) read(reader *ClassReader) {
	this.ConstantMemberRefInfo.read(reader)
}

// --------------------- 接口方法引用 -------------------------

type ConstantInterfaceMethodRefInfo struct {
	ConstantMemberRefInfo
}

func (this *ConstantInterfaceMethodRefInfo) Val() interface{} {
	panic("implement me")
}

func (this *ConstantInterfaceMethodRefInfo) String() string {
	return fmt.Sprintf("{ConstantInterfaceMethodRefInfo: %v}", this.ConstantMemberRefInfo.String())
}

func (this *ConstantInterfaceMethodRefInfo) read(reader *ClassReader) {
	this.ConstantMemberRefInfo.read(reader)
}

// todo 动态调用相关常量解析
