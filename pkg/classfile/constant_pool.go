package classfile

import (
	"fmt"
	"strconv"
	"strings"

	"jvm/pkg/global"
)

const (
	CONSTANT_Class              = 7
	CONSTANT_Fieldref           = 9
	CONSTANT_Methodref          = 10
	CONSTANT_InterfaceMethodref = 11
	CONSTANT_String             = 8
	CONSTANT_Integer            = 3
	CONSTANT_Float              = 4
	CONSTANT_Long               = 5
	CONSTANT_Double             = 6
	CONSTANT_NameAndType        = 12
	CONSTANT_Utf8               = 1
	CONSTANT_MethodHandle       = 15
	CONSTANT_MethodType         = 16
	CONSTANT_InvokeDynamic      = 18
)

type ConstantPool []ConstantInfo

type ConstantInfo interface {
	read(reader *ClassReader)

	Val() interface{}
}

func readConstantPool(reader *ClassReader) ConstantPool {
	count := int(reader.readUnit16())
	constantPool := make([]ConstantInfo, count)
	if global.Verbose {
		log.Infof("parsed constant pool length : %v", count)
	}

	for i := 1; i < count; i++ {
		constantPool[i] = readConstantInfo(reader, constantPool)
		switch constantPool[i].(type) {
		case *ConstantDoubleInfo, *ConstantLongInfo:
			i++
		}
	}

	if global.Verbose {
		log.Infof("parsed constant pool: %v", constantPool)
	}

	return constantPool
}

func readConstantInfo(reader *ClassReader, cp ConstantPool) ConstantInfo {
	tag := reader.readUint8()
	constantInfo := newConstantInfo(tag, cp)

	constantInfo.read(reader)
	return constantInfo
}

func newConstantInfo(tag uint8, cp ConstantPool) ConstantInfo {
	switch tag {
	case CONSTANT_Integer:
		return &ConstantIntegerInfo{}
	case CONSTANT_Long:
		return &ConstantLongInfo{}
	case CONSTANT_Float:
		return &ConstantFloatInfo{}
	case CONSTANT_Double:
		return &ConstantDoubleInfo{}
	case CONSTANT_Utf8:
		return &ConstantUtf8Info{}
	case CONSTANT_Class:
		return &ConstantClassInfo{cp: cp}
	case CONSTANT_String:
		return &ConstantStringInfo{cp: cp}
	case CONSTANT_Fieldref:
		return &ConstantFieldRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_Methodref:
		return &ConstantMethodRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_InterfaceMethodref:
		return &ConstantInterfaceMethodRefInfo{ConstantMemberRefInfo{
			cp: cp,
		}}
	case CONSTANT_NameAndType:
		return &ConstantNameAndTypeRefInfo{
			cp: cp,
		}
	case CONSTANT_MethodHandle:
		return &ConstantMethodHandleInfo{}
	case CONSTANT_MethodType:
		return &ConstantMethodTypeInfo{}
	case CONSTANT_InvokeDynamic:
		return &ConstantInvokeDynamicInfo{}
	default:
		log.Panicf("error constant info tag: %v", tag)
		return nil
	}
}

func (this ConstantPool) String() string {
	builder := strings.Builder{}
	builder.WriteString("ConstantPool{\n")
	length := len(this)
	for i := 1; i < length; i++ {
		constantInfo := this[i]

		builder.WriteString(strings.Join([]string{"#", strconv.Itoa(i), " ", fmt.Sprintf("%v", constantInfo), "\n"}, ""))

		switch constantInfo.(type) {
		case *ConstantDoubleInfo, *ConstantLongInfo:
			i++
		}
	}
	builder.WriteString("}")

	return builder.String()
}
