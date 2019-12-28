package global

import "os"

//  -------------------------------     变量区     ----------------------------------

var Verbose = false

//  -------------------------------     常量区     ----------------------------------

const Gava string = "gava"
const Version string = "version 0.0.1"

const PathListSeparator = string(os.PathListSeparator)
const Space string = " "
const EmptyString string = ""
const Dot string = "."
const Slash string = "/"
const WildCard string = "*"
const Semicolon string = ";"

const SuffixJar string = ".jar"
const SuffixZip string = ".zip"
const SuffixClass string = ".class"

const JavaHome string = "JAVA_HOME"

const JavaClassFileMagic uint32 = 0xCAFEBABE

const KeywordPublic string = "public"
const KeywordProtected string = "protected"
const KeywordPrivate string = "private"
const KeywordAbstract string = "abstract"
const KeywordFinal string = "final"
const KeywordBridge string = "bridge"
const KeywordStrict string = "strict"
const KeywordNative string = "native"
const KeywordSynchronized string = "synchronized"
const KeywordVolatile string = "volatile"
const KeywordTransient string = "transient"
const KeywordInterface string = "interface"
const KeywordAnnotation string = "@interface"
const KeywordEnum string = "enum"
const KeywordStatic string = "static"
const AccGenerated string = "generated"

const Main string = "main"
const MainDescriptor string = "([Ljava/lang/String;)V"

const FdBoolean string = "Z"
const FdByte string = "B"
const FdChar string = "C"
const FdShort string = "S"
const FdInt string = "I"
const FdLong string = "J"
const FdFloat string = "F"
const FdDouble string = "D"
const FdString string = "Ljava/lang/String;"
const FdRef string = "L"
const FdArray string = "["
const JavaLangObject string = "java/lang/Object"
const JavaLangCloneable string = "java/lang/Cloneable"
const JavaIOSerializable string = "java/io/Serializable"
const JavaLangClass string = "java/lang/Class"
const JavaLangString string = "java/lang/String"

const (
	OpcNop uint8 = iota // NOP 对应 0x00, 后面自增1
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
