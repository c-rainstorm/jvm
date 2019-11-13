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

const SuffixJar string = ".jar"
const SuffixZip string = ".zip"
const SuffixClass string = ".class"

const JavaHome string = "JAVA_HOME"

const JavaClassFileMagic uint32 = 0xCAFEBABE
