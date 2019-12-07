package heap

import "jvm/pkg/global"

type Field struct {
	ClassMember
	// 字段对应的槽号
	slotId uint
	// 常量值下标，当当前字段没有常量初始值时下标为 0
	constValueIndex uint
}

func (this *Field) isStatic() bool {
	return this.hasFlag(ACC_STATIC)
}

func (this *Field) isDoubleOrLong() bool {
	return this.descriptor == global.FdDouble || this.descriptor == global.FdLong
}

func (this *Field) isFinal() bool {
	return this.hasFlag(ACC_FINAL)
}

func (this *Field) ConstValueIndex() uint {
	return this.constValueIndex
}
