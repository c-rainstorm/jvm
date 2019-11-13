package classpath

import (
	"strings"

	"jvm/pkg/global"
)

type Entry interface {
	// classname 参数是相对路径，比如
	// 		java.lang.Object -> java/lang/Object.class
	readClass(classname string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, global.PathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, global.WildCard) {
		return newWildCardEntry(path)
	}

	lowerCasePath := strings.ToLower(path)
	if strings.HasSuffix(lowerCasePath, global.SuffixJar) ||
		strings.HasSuffix(lowerCasePath, global.SuffixZip) {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
