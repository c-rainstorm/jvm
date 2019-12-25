package heap

import (
	"unicode/utf16"

	"jvm/pkg/global"
)

var internedStrings = map[string]*NormalObject{}

func JString(loader *ClassLoader, goStr string) *NormalObject {
	if internedStr, ok := internedStrings[goStr]; ok {
		return internedStr
	}

	chars := stringToUTF16(goStr)
	jChars := &ArrayObject{BaseObject: BaseObject{class: loader.LoadClass(global.FdArray + global.FdChar)}, data: chars}

	jStr := loader.LoadClass(global.JavaLangString).NewObject().(*NormalObject)
	jStr.Class().SetField(jStr, "value", global.FdArray+global.FdChar, jChars)

	internedStrings[goStr] = jStr
	return jStr
}

func stringToUTF16(str string) []uint16 {
	return utf16.Encode([]rune(str))
}

func UTF16ToString(s [] uint16) string {
	return string(utf16.Decode(s))
}

func GoString(jStr *NormalObject) string {
	jChars := jStr.Class().GetField(jStr, "value", global.FdArray+global.FdChar).(*ArrayObject)
	return UTF16ToString(jChars.data.([]uint16))
}

func Intern0(jStr *NormalObject) Object {
	goStr := GoString(jStr)
	if interned, ok := internedStrings[goStr]; ok {
		return interned
	}

	internedStrings[goStr] = jStr
	return jStr
}
