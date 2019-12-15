package heap

import "jvm/pkg/global"

type Field struct {
	ClassMember
	// 字段对应的槽号
	slotId uint
	// 常量值下标，当当前字段没有常量初始值时下标为 0
	constValueIndex uint
}

func (this *Field) IsStatic() bool {
	return this.hasFlag(ACC_STATIC)
}

func (this *Field) IsDoubleOrLong() bool {
	return this.descriptor == global.FdDouble || this.descriptor == global.FdLong
}

func (this *Field) IsFinal() bool {
	return this.hasFlag(ACC_FINAL)
}

func (this *Field) ConstValueIndex() uint {
	return this.constValueIndex
}

func (this *Field) SlotId() uint {
	return this.slotId
}
