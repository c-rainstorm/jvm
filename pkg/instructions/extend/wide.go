package extend

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/instructions/loads"
	"jvm/pkg/instructions/math"
	"jvm/pkg/instructions/stores"
	"jvm/pkg/rtda"
)

const (
	WILoad  uint8 = 0x15
	WLLoad  uint8 = 0x16
	WFLoad  uint8 = 0x17
	WDLoad  uint8 = 0x18
	WALoad  uint8 = 0x19
	WIStore uint8 = 0x36
	WLStore uint8 = 0x37
	WFStore uint8 = 0x38
	WDStore uint8 = 0x39
	WAStore uint8 = 0x3a
	WIInc   uint8 = 0x84
	WRet    uint8 = 0xa9
)

type Wide struct {
	modifiedInstruction base.Instruction
}

func (this *Wide) FetchOperands(reader *base.ByteCodeReader) {
	opCode := reader.ReadUint8()
	switch opCode {
	case WILoad:
		inst := &loads.ILoad{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WLLoad:
		inst := &loads.LLoad{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WFLoad:
		inst := &loads.FLoad{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WDLoad:
		inst := &loads.DLoad{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WALoad:
		inst := &loads.ALoad{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WIStore:
		inst := &stores.IStore{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WLStore:
		inst := &stores.LStore{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WFStore:
		inst := &stores.FStore{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WDStore:
		inst := &stores.DStore{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WAStore:
		inst := &stores.AStore{}
		inst.Index = uint(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WIInc:
		inst := &math.IInc{}
		inst.Index = uint(reader.ReadUint16())
		inst.Const = int32(reader.ReadUint16())
		this.modifiedInstruction = inst
	case WRet:
		panic("Unsupported opcode : ret(0xa9)")
	}
}

func (this *Wide) Execute(frame *rtda.Frame) {
	this.modifiedInstruction.Execute(frame)
}
