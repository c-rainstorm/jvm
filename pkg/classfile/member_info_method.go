package classfile

import (
	"fmt"
	"strings"

	"jvm/pkg/global"
)

const (
	accMethodPublic       uint16 = 0x0001
	accMethodPrivate      uint16 = 0x0002
	accMethodProtected    uint16 = 0x0004
	accMethodStatic       uint16 = 0x0008
	accMethodFinal        uint16 = 0x0010
	accMethodSynchronized uint16 = 0x0020
	accMethodBridge       uint16 = 0x0040
	accMethodVarargs      uint16 = 0x0080
	accMethodNative       uint16 = 0x0100
	accMethodAbstract     uint16 = 0x0400
	accMethodStrict       uint16 = 0x0800
	accMethodSynthetic    uint16 = 0x1000
)

type MethodMemberInfo struct {
	MemberInfo
}

func (this *MethodMemberInfo) read(reader *ClassReader, cp ConstantPool) {
	this.MemberInfo.read(reader, cp)
}

func (this *MethodMemberInfo) AccessFlags() string {
	builder := strings.Builder{}
	builder.WriteString(this.checkAccessFlag(accMethodPublic, global.KeywordPublic))
	builder.WriteString(this.checkAccessFlag(accMethodProtected, global.KeywordProtected))
	builder.WriteString(this.checkAccessFlag(accMethodPrivate, global.KeywordPrivate))
	builder.WriteString(this.checkAccessFlag(accMethodStatic, global.KeywordStatic))
	builder.WriteString(this.checkAccessFlag(accMethodFinal, global.KeywordFinal))
	builder.WriteString(this.checkAccessFlag(accMethodAbstract, global.KeywordAbstract))
	// 声明同步方法
	builder.WriteString(this.checkAccessFlag(accMethodSynchronized, global.KeywordSynchronized))
	// 使用非 Java 语言实现的方法
	builder.WriteString(this.checkAccessFlag(accMethodNative, global.KeywordNative))
	builder.WriteString(this.checkAccessFlag(accMethodStrict, global.KeywordStrict))
	// 桥接方法，由编译器生成，泛型方法继承实现时会由编译器自动生成一个桥接方法
	//
	// public class compareValues implements Comparator {
	//
	//   // target method
	//   public int compare(String a, String b)
	//   {
	//   }
	//
	//   // bridge method
	//   public int compare(Object a, Object b) {
	//      return compare((String)a, (String)b);
	//   }
	// }
	//  https://www.geeksforgeeks.org/method-class-isbridge-method-in-java/
	builder.WriteString(this.checkAccessFlag(accMethodBridge, global.KeywordBridge))
	builder.WriteString(this.checkAccessFlag(accMethodSynthetic, global.AccGenerated))
	return builder.String()
}

func (this *MethodMemberInfo) checkAccessFlag(targetFlag uint16, targetKeyword string) string {
	return this.MemberInfo.checkAccessFlag(targetFlag, targetKeyword)
}

func (this *MethodMemberInfo) String() string {
	return fmt.Sprintf("%s %s %s", this.AccessFlags(), this.cp[this.descriptorIndex].Val(), this.cp[this.nameIndex].Val())
}
