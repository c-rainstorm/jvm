package heap

import (
	"unicode/utf16"

	"jvm/pkg/global"
)

var internedStrings = map[string]*Object{}

func JString(loader *ClassLoader, goStr string) *Object {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUTF16(goStr)
	jChars := &Object{class: loader.LoadClass(global.FdArray + global.FdChar), data: chars}

	jStr := loader.LoadClass("java/lang/String").NewObject()
	jStr.SetField("value", global.FdArray+global.FdChar, jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUTF16(str string) []uint16 {
	return utf16.Encode([]rune(str))
}

func UTF16ToString(s [] uint16) string {
	return string(utf16.Decode(s))
}

func GoString(jStr *Object) string {
	jChars := jStr.GetField("value", global.FdArray+global.FdChar).(*Object)
	return UTF16ToString(jChars.data.([]uint16))
}
