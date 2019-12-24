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
