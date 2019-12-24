package heap

import (
	"jvm/pkg/classfile"
)

type ClassSymRef struct {
	SymbolicRef
}

func newClassSymRef(cp *ConstantPool, cfClass *classfile.ConstantClassInfo) Constant {
	classSymRef := &ClassSymRef{}
	classSymRef.cp = cp
	classSymRef.className = cfClass.Name()
	return classSymRef
}
