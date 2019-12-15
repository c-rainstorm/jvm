package instructions

import (
	"log"

	"jvm/pkg/instructions/base"
	"jvm/pkg/instructions/comparisons"
	"jvm/pkg/instructions/constants"
	"jvm/pkg/instructions/control"
	"jvm/pkg/instructions/conversions"
	"jvm/pkg/instructions/extend"
	"jvm/pkg/instructions/loads"
	"jvm/pkg/instructions/math"
	"jvm/pkg/instructions/references"
	"jvm/pkg/instructions/stack"
	"jvm/pkg/instructions/stores"
)

// IDEA 行号减20即为字节码对应数字  NOP 对应 0x00
const (
	OpcNop uint8 = iota
	OpcAConstNull
	OpcIConstM1
	OpcIConst0
	OpcIConst1
	OpcIConst2
	OpcIConst3
	OpcIConst4
	OpcIConst5
	OpcLConst0
	OpcLConst1
	OpcFConst0
	OpcFConst1
	OpcFConst2
	OpcDConst0
	OpcDConst1
	OpcBIPush
	OpcSIPush
	OpcLDC
	OpcLDCW
	OpcLDC2W
	OpcILoad
	OpcLLoad
	OpcFLoad
	OpcDLoad
	OpcALoad
	OpcILoad0
	OpcILoad1
	OpcILoad2
	OpcILoad3
	OpcLLoad0
	OpcLLoad1
	OpcLLoad2
	OpcLLoad3
	OpcFLoad0
	OpcFLoad1
	OpcFLoad2
	OpcFLoad3
	OpcDLoad0
	OpcDLoad1
	OpcDLoad2
	OpcDLoad3
	OpcALoad0
	OpcALoad1
	OpcALoad2
	OpcALoad3
	OpcIALoad
	OpcLALoad
	OpcFALoad
	OpcDALoad
	OpcAALoad
	OpcBALoad
	OpcCALoad
	OpcSALoad
	OpcIStore
	OpcLStore
	OpcFStore
	OpcDStore
	OpcAStore
	OpcIStore0
	OpcIStore1
	OpcIStore2
	OpcIStore3
	OpcLStore0
	OpcLStore1
	OpcLStore2
	OpcLStore3
	OpcFStore0
	OpcFStore1
	OpcFStore2
	OpcFStore3
	OpcDStore0
	OpcDStore1
	OpcDStore2
	OpcDStore3
	OpcAStore0
	OpcAStore1
	OpcAStore2
	OpcAStore3
	OpcIAStore
	OpcLAStore
	OpcFAStore
	OpcDAStore
	OpcAAStore
	OpcBAStore
	OpcCAStore
	OpcSAStore
	OpcPOP
	OpcPOP2
	OpcDup
	OpcDupX1
	OpcDupX2
	OpcDup2
	OpcDup2X1
	OpcDup2X2
	OpcSwap
	OpcIAdd
	OpcLAdd
	OpcFAdd
	OpcDAdd
	OpcISub
	OpcLSub
	OpcFSub
	OpcDSub
	OpcIMul
	OpcLMul
	OpcFMul
	OpcDMul
	OpcIDiv
	OpcLDiv
	OpcFDiv
	OpcDDiv
	OpcIRem
	OpcLRem
	OpcFRem
	OpcDRem
	OpcINeg
	OpcLNeg
	OpcFNeg
	OpcDNeg
	OpcIShL
	OpcLShL
	OpcIShR
	OpcLShR
	OpcIUShR
	OpcLUShR
	OpcIAnd
	OpcLAnd
	OpcIOr
	OpcLOr
	OpcIXOr
	OpcLXOr
	OpcIInc
	OpcI2L
	OpcI2F
	OpcI2D
	OpcL2I
	OpcL2F
	OpcL2D
	OpcF2I
	OpcF2L
	OpcF2D
	OpcD2I
	OpcD2L
	OpcD2F
	OpcI2B
	OpcI2C
	OpcI2S
	OpcLCmp
	OpcFCmpL
	OpcFCmpG
	OpcDCmpL
	OpcDCmpG
	OpcIfEQ
	OpcIfNE
	OpcIfLT
	OpcIfGE
	OpcIfGT
	OpcIfLE
	OpcIfICmpEQ
	OpcIfICmpNE
	OpcIfICmpLT
	OpcIfICmpGE
	OpcIfICmpGT
	OpcIfICmpLE
	OpcIfACmpEQ
	OpcIfACmpNE
	OpcGOTO
	OpcJSR
	OpcRet
	OpcTableSwitch
	OpcLookupSwitch
	OpcIReturn
	OpcLReturn
	OpcFReturn
	OpcDReturn
	OpcAReturn
	OpcReturn
	OpcGetStatic
	OpcPutStatic
	OpcGetField
	OpcPutField
	OpcInvokeVirtual
	OpcInvokeSpecial
	OpcInvokeStatic
	OpcInvokeInterface
	OpcInvokeDynamic
	OpcNew
	OpcNewArray
	OpcANewArray
	OpcArrayLength
	OpcAThrow
	OpcCheckCast
	OpcInstanceOf
	OpcMonitorEnter
	OpcMonitorExit
	OpcWide
	OpcMultiANewArray
	OpcIfNull
	OpcIfNonNull
	OpcGotoW
	OpcJsrW
	OpcBreakPoint
	OpcImpDep1 = 0xFE
	OpcImpDep2 = 0xFF
)

func New(opCode uint8) base.Instruction {
	switch opCode {
	case OpcNop:
		return &constants.NOP{}
	case OpcAConstNull:
		return &constants.AConstNull{}
	case OpcIConstM1:
		return &constants.IConstM1{}
	case OpcIConst0:
		return &constants.IConst0{}
	case OpcIConst1:
		return &constants.IConst1{}
	case OpcIConst2:
		return &constants.IConst2{}
	case OpcIConst3:
		return &constants.IConst3{}
	case OpcIConst4:
		return &constants.IConst4{}
	case OpcIConst5:
		return &constants.IConst5{}
	case OpcLConst0:
		return &constants.LConst0{}
	case OpcLConst1:
		return &constants.LConst1{}
	case OpcFConst0:
		return &constants.FConst0{}
	case OpcFConst1:
		return &constants.FConst1{}
	case OpcFConst2:
		return &constants.FConst2{}
	case OpcDConst0:
		return &constants.DConst0{}
	case OpcDConst1:
		return &constants.DConst1{}
	case OpcBIPush:
		return &constants.BIPush{}
	case OpcSIPush:
		return &constants.SIPush{}
	case OpcLDC:
		return &constants.LDC{}
	case OpcLDCW:
		return &constants.LDCW{}
	case OpcLDC2W:
		return &constants.LDC2W{}
	case OpcILoad:
		return &loads.ILoad{}
	case OpcLLoad:
		return &loads.LLoad{}
	case OpcFLoad:
		return &loads.FLoad{}
	case OpcDLoad:
		return &loads.DLoad{}
	case OpcALoad:
		return &loads.ALoad{}
	case OpcILoad0:
		return &loads.ILoad0{}
	case OpcILoad1:
		return &loads.ILoad1{}
	case OpcILoad2:
		return &loads.ILoad2{}
	case OpcILoad3:
		return &loads.ILoad3{}
	case OpcLLoad0:
		return &loads.LLoad0{}
	case OpcLLoad1:
		return &loads.LLoad1{}
	case OpcLLoad2:
		return &loads.LLoad2{}
	case OpcLLoad3:
		return &loads.LLoad3{}
	case OpcFLoad0:
		return &loads.FLoad0{}
	case OpcFLoad1:
		return &loads.FLoad1{}
	case OpcFLoad2:
		return &loads.FLoad2{}
	case OpcFLoad3:
		return &loads.FLoad3{}
	case OpcDLoad0:
		return &loads.DLoad0{}
	case OpcDLoad1:
		return &loads.DLoad1{}
	case OpcDLoad2:
		return &loads.DLoad2{}
	case OpcDLoad3:
		return &loads.DLoad3{}
	case OpcALoad0:
		return &loads.ALoad0{}
	case OpcALoad1:
		return &loads.ALoad1{}
	case OpcALoad2:
		return &loads.ALoad2{}
	case OpcALoad3:
		return &loads.ALoad3{}
	case OpcIALoad:
		return nil
	case OpcLALoad:
		return nil
	case OpcFALoad:
		return nil
	case OpcDALoad:
		return nil
	case OpcAALoad:
		return nil
	case OpcBALoad:
		return nil
	case OpcCALoad:
		return nil
	case OpcSALoad:
		return nil
	case OpcIStore:
		return &stores.IStore{}
	case OpcLStore:
		return &stores.LStore{}
	case OpcFStore:
		return &stores.FStore{}
	case OpcDStore:
		return &stores.DStore{}
	case OpcAStore:
		return &stores.AStore{}
	case OpcIStore0:
		return &stores.IStore0{}
	case OpcIStore1:
		return &stores.IStore1{}
	case OpcIStore2:
		return &stores.IStore2{}
	case OpcIStore3:
		return &stores.IStore3{}
	case OpcLStore0:
		return &stores.LStore0{}
	case OpcLStore1:
		return &stores.LStore1{}
	case OpcLStore2:
		return &stores.LStore2{}
	case OpcLStore3:
		return &stores.LStore3{}
	case OpcFStore0:
		return &stores.FStore0{}
	case OpcFStore1:
		return &stores.FStore1{}
	case OpcFStore2:
		return &stores.FStore2{}
	case OpcFStore3:
		return &stores.FStore3{}
	case OpcDStore0:
		return &stores.DStore0{}
	case OpcDStore1:
		return &stores.DStore1{}
	case OpcDStore2:
		return &stores.DStore2{}
	case OpcDStore3:
		return &stores.DStore3{}
	case OpcAStore0:
		return &stores.AStore0{}
	case OpcAStore1:
		return &stores.AStore1{}
	case OpcAStore2:
		return &stores.AStore2{}
	case OpcAStore3:
		return &stores.AStore3{}
	case OpcIAStore:
		return nil
	case OpcLAStore:
		return nil
	case OpcFAStore:
		return nil
	case OpcDAStore:
		return nil
	case OpcAAStore:
		return nil
	case OpcBAStore:
		return nil
	case OpcCAStore:
		return nil
	case OpcSAStore:
		return nil
	case OpcPOP:
		return &stack.POP{}
	case OpcPOP2:
		return &stack.POP2{}
	case OpcDup:
		return &stack.Dup{}
	case OpcDupX1:
		return &stack.DupX1{}
	case OpcDupX2:
		return &stack.DupX2{}
	case OpcDup2:
		return &stack.Dup2{}
	case OpcDup2X1:
		return &stack.Dup2X1{}
	case OpcDup2X2:
		return &stack.Dup2X2{}
	case OpcSwap:
		return &stack.Swap{}
	case OpcIAdd:
		return &math.IAdd{}
	case OpcLAdd:
		return &math.LAdd{}
	case OpcFAdd:
		return &math.FAdd{}
	case OpcDAdd:
		return &math.DAdd{}
	case OpcISub:
		return &math.ISub{}
	case OpcLSub:
		return &math.LSub{}
	case OpcFSub:
		return &math.FSub{}
	case OpcDSub:
		return &math.DSub{}
	case OpcIMul:
		return &math.IMul{}
	case OpcLMul:
		return &math.LMul{}
	case OpcFMul:
		return &math.FMul{}
	case OpcDMul:
		return &math.DMul{}
	case OpcIDiv:
		return &math.IDiv{}
	case OpcLDiv:
		return &math.LDiv{}
	case OpcFDiv:
		return &math.FDiv{}
	case OpcDDiv:
		return &math.DDiv{}
	case OpcIRem:
		return &math.IRem{}
	case OpcLRem:
		return &math.LRem{}
	case OpcFRem:
		return &math.FRem{}
	case OpcDRem:
		return &math.DRem{}
	case OpcINeg:
		return &math.INeg{}
	case OpcLNeg:
		return &math.LNeg{}
	case OpcFNeg:
		return &math.FNeg{}
	case OpcDNeg:
		return &math.DNeg{}
	case OpcIShL:
		return &math.IShL{}
	case OpcLShL:
		return &math.LShL{}
	case OpcIShR:
		return &math.IShR{}
	case OpcLShR:
		return &math.LShR{}
	case OpcIUShR:
		return &math.IUShR{}
	case OpcLUShR:
		return &math.LUShR{}
	case OpcIAnd:
		return &math.IAnd{}
	case OpcLAnd:
		return &math.LAnd{}
	case OpcIOr:
		return &math.IOr{}
	case OpcLOr:
		return &math.LOr{}
	case OpcIXOr:
		return &math.IXOr{}
	case OpcLXOr:
		return &math.LXOr{}
	case OpcIInc:
		return &math.IInc{}
	case OpcI2L:
		return &conversions.I2F{}
	case OpcI2F:
		return &conversions.I2F{}
	case OpcI2D:
		return &conversions.I2D{}
	case OpcL2I:
		return &conversions.L2I{}
	case OpcL2F:
		return &conversions.L2F{}
	case OpcL2D:
		return &conversions.L2D{}
	case OpcF2I:
		return &conversions.F2I{}
	case OpcF2L:
		return &conversions.F2L{}
	case OpcF2D:
		return &conversions.F2D{}
	case OpcD2I:
		return &conversions.D2L{}
	case OpcD2L:
		return &conversions.D2L{}
	case OpcD2F:
		return &conversions.D2F{}
	case OpcI2B:
		return &conversions.I2B{}
	case OpcI2C:
		return &conversions.I2C{}
	case OpcI2S:
		return &conversions.I2S{}
	case OpcLCmp:
		return &comparisons.LCmp{}
	case OpcFCmpL:
		return &comparisons.FCmpL{}
	case OpcFCmpG:
		return &comparisons.FCmpG{}
	case OpcDCmpL:
		return &comparisons.DCmpL{}
	case OpcDCmpG:
		return &comparisons.DCmpG{}
	case OpcIfEQ:
		return &comparisons.IfEQ{}
	case OpcIfNE:
		return &comparisons.IfNE{}
	case OpcIfLT:
		return &comparisons.IfLT{}
	case OpcIfGE:
		return &comparisons.IfGE{}
	case OpcIfGT:
		return &comparisons.IfGT{}
	case OpcIfLE:
		return &comparisons.IfLE{}
	case OpcIfICmpEQ:
		return &comparisons.IfICmpEQ{}
	case OpcIfICmpNE:
		return &comparisons.IfICmpNE{}
	case OpcIfICmpLT:
		return &comparisons.IfICmpLT{}
	case OpcIfICmpGE:
		return &comparisons.IfICmpGE{}
	case OpcIfICmpGT:
		return &comparisons.IfICmpGT{}
	case OpcIfICmpLE:
		return &comparisons.IfICmpLE{}
	case OpcIfACmpEQ:
		return &comparisons.IfACmpEQ{}
	case OpcIfACmpNE:
		return &comparisons.IfACmpNE{}
	case OpcGOTO:
		return &control.GOTO{}
	case OpcJSR:
		return nil
	case OpcRet:
		log.Panic("ret not found")
		return nil
	case OpcTableSwitch:
		return &control.TableSwitch{}
	case OpcLookupSwitch:
		return &control.LookupSwitch{}
	case OpcIReturn:
		return &control.IReturn{}
	case OpcLReturn:
		return &control.LReturn{}
	case OpcFReturn:
		return &control.FReturn{}
	case OpcDReturn:
		return &control.DReturn{}
	case OpcAReturn:
		return &control.AReturn{}
	case OpcReturn:
		return &control.Return{}
	case OpcGetStatic:
		return &references.GetStatic{}
	case OpcPutStatic:
		return &references.PutStatic{}
	case OpcGetField:
		return &references.GetField{}
	case OpcPutField:
		return &references.PutField{}
	case OpcInvokeVirtual:
		return &references.InvokeVirtual{}
	case OpcInvokeSpecial:
		return &references.InvokeSpecial{}
	case OpcInvokeStatic:
		return &references.InvokeStatic{}
	case OpcInvokeInterface:
		return &references.InvokeInterface{}
	case OpcInvokeDynamic:
		return nil
	case OpcNew:
		return &references.New{}
	case OpcNewArray:
		return nil
	case OpcANewArray:
		return nil
	case OpcArrayLength:
		return nil
	case OpcAThrow:
		return nil
	case OpcCheckCast:
		return &references.CheckCast{}
	case OpcInstanceOf:
		return &references.InstanceOf{}
	case OpcMonitorEnter:
		return nil
	case OpcMonitorExit:
		return nil
	case OpcWide:
		return &extend.Wide{}
	case OpcMultiANewArray:
		return nil
	case OpcIfNull:
		return &extend.IfNull{}
	case OpcIfNonNull:
		return &extend.IfNonNull{}
	case OpcGotoW:
		return &extend.GOTOW{}
	case OpcJsrW:
		return nil
	case OpcBreakPoint:
		return nil
	case OpcImpDep1:
		return nil
	case OpcImpDep2:
		return nil
	default:
		log.Panicf("当前指令未实现：%X", opCode)
		return nil
	}
	return nil
}
