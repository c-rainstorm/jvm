package classfile

import (
	"fmt"
	"strings"

	"jvm/pkg/global"
)

const (
	accFieldPublic    uint16 = 0x0001
	accFieldPrivate   uint16 = 0x0002
	accFieldProtected uint16 = 0x0004
	accFieldStatic    uint16 = 0x0008
	accFieldFinal     uint16 = 0x0010
	accFieldVolatile  uint16 = 0x0040
	accFieldTransient uint16 = 0x0080
	accFieldSynthetic uint16 = 0x1000
	accFieldEnum      uint16 = 0x4000
)

type FieldMemberInfo struct {
	MemberInfo
}

func (this *FieldMemberInfo) read(reader *ClassReader, cp ConstantPool) {
	this.MemberInfo.read(reader, cp)
}

func (this *FieldMemberInfo) AccessFlags() string {
	builder := strings.Builder{}
	builder.WriteString(this.checkAccessFlag(accFieldPublic, global.KeywordPublic))
	builder.WriteString(this.checkAccessFlag(accFieldProtected, global.KeywordProtected))
	builder.WriteString(this.checkAccessFlag(accFieldPrivate, global.KeywordPrivate))
	builder.WriteString(this.checkAccessFlag(accFieldStatic, global.KeywordStatic))
	builder.WriteString(this.checkAccessFlag(accFieldFinal, global.KeywordFinal))
	builder.WriteString(this.checkAccessFlag(accFieldVolatile, global.KeywordVolatile))
	builder.WriteString(this.checkAccessFlag(accFieldTransient, global.KeywordTransient))
	builder.WriteString(this.checkAccessFlag(accFieldEnum, global.KeywordEnum))
	builder.WriteString(this.checkAccessFlag(accFieldSynthetic, global.AccGenerated))
	return builder.String()
}

func (this *FieldMemberInfo) checkAccessFlag(targetFlag uint16, targetKeyword string) string {
	return this.MemberInfo.checkAccessFlag(targetFlag, targetKeyword)
}

func (this *FieldMemberInfo) String() string {
	return fmt.Sprintf("%s %s %s", this.AccessFlags(), this.cp[this.descriptorIndex].Val(), this.cp[this.nameIndex].Val())
}
