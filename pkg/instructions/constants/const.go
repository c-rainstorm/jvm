package constants

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// aconst_null   Page.370
//
// 格式: aconst_null
// 字节: 0x1
// 操作: 将 null 对象引用推到操作数栈
// 描述: JVM 并不限定 null 的具体值
//
type AConstNull struct {
	base.NoOperandsInstruction
}

func (this *AConstNull) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushRef(nil)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// iconst_<i>   Page.456
//
// 格式: iconst_<i>
// 字节:
//   助记符      8进制表示   对应整数
//   iconst_m1    0x2       -1
//   iconst_0     0x3       0
//   iconst_1     0x4       1
//   iconst_2     0x5       2
//   iconst_3     0x6       3
//   iconst_4     0x7       4
//   iconst_5     0x8       5
// 操作: 将整型i推到操作数栈
// 描述: 该指令族和 bipush [byte num] 指令的效果相同，只不过该指令族在指令中隐含了推到操作数栈的整数
//
type IConstM1 struct {
	base.NoOperandsInstruction
}

func (this *IConstM1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(-1)
}

type IConst0 struct {
	base.NoOperandsInstruction
}

func (this *IConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(0)
}

type IConst1 struct {
	base.NoOperandsInstruction
}

func (this *IConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(1)
}


type IConst2 struct {
	base.NoOperandsInstruction
}

func (this *IConst2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(2)
}

type IConst3 struct {
	base.NoOperandsInstruction
}

func (this *IConst3) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(3)
}

type IConst4 struct {
	base.NoOperandsInstruction
}

func (this *IConst4) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(4)
}

type IConst5 struct {
	base.NoOperandsInstruction
}

func (this *IConst5) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(5)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// lconst_<i>   Page.514
//
// 格式: lconst_<i>
// 字节:
//   助记符      8进制表示   对应整数
//   lconst_0     0x9       0
//   lconst_1     0xa       1
// 操作: 将long型i推到操作数栈
// 描述:
//
type LConst0 struct {
	base.NoOperandsInstruction
}

func (this *LConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(0)
}

type LConst1 struct {
	base.NoOperandsInstruction
}

func (this *LConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushLong(1)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// fconst_<i>   Page.426
//
// 格式: fconst_<i>
// 字节:
//   助记符      8进制表示   对应浮点数
//   fconst_0     0xb       0.0
//   fconst_1     0xc       1.0
//   fconst_2     0xd       2.0
// 操作: 将float型(0.0, or 1.0, or 2.0)推到操作数栈
// 描述:
//
type FConst0 struct {
	base.NoOperandsInstruction
}

func (this *FConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(0)
}

type FConst1 struct {
	base.NoOperandsInstruction
}

func (this *FConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(1)
}

type FConst2 struct {
	base.NoOperandsInstruction
}

func (this *FConst2) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushFloat(2)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// dconst_<i>   Page.396
//
// 格式: dconst_<i>
// 字节:
//   助记符      8进制表示   对应双进度浮点数
//   dconst_0     0xe       0.0
//   dconst_1     0xf       1.0
// 操作: 将double型(0.0 or 1.0)推到操作数栈
// 描述:
//
type DConst0 struct {
	base.NoOperandsInstruction
}

func (this *DConst0) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(0.0)
}

type DConst1 struct {
	base.NoOperandsInstruction
}

func (this *DConst1) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushDouble(1.0)
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// bipush   Page.396
//
// 格式: bipush [byte]
// 字节: 0x10 0x1     //
//   指令样例      8进制表示    指令含义
//   bipush 1     0x10 0x01   将 1 转成整数推到操作数栈 (0x01 -> 0x00000001)
//   bipush -1    0x10 0xFF   将 -1 转成整数推到操作数栈 (0xFF -> 0xFFFFFFFF)
// 操作: 将指令后面的一字节的立即数【符号拓展(带符号)】为 int 型，推到操作数栈
// 描述:
//
type BIPush struct {
	val int8
}

func (this *BIPush) FetchOperands(reader *base.ByteCodeReader) {
	this.val = int8(reader.ReadUint8())
}

func (this *BIPush) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(this.val))
}

// 指令定义来源 The Java® Virtual Machine Specification
// https://docs.oracle.com/javase/specs/jvms/se8/jvms8.pdf
// sipush   Page.558
//
// 格式: sipush [byte1] [byte2]
// 字节: 0x11 0x00 0x01     //
//   指令样例      8进制表示         指令含义
//   sipush 1     0x11 0x00 0x01   将 1 转成整数推到操作数栈 (0x0001 -> 0x00000001)
//   sipush -1    0x11 0xFF 0xFF   将 -1 转成整数推到操作数栈 (0xFFFF -> 0xFFFFFFFF)
// 操作: 将指令后面的两字节的立即数【符号拓展(带符号)】为 int 型，推到操作数栈
// 描述:
//
type SIPush struct {
	val int16
}

func (this *SIPush) FetchOperands(reader *base.ByteCodeReader) {
	this.val = int16(reader.ReadUint16())
}

func (this *SIPush) Execute(frame *rtda.Frame) {
	frame.OperandStack().PushInt(int32(this.val))
}
