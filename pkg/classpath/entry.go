package classpath

import (
	"jvm/pkg/constants"
	"strings"
)

type Entry interface {
	// classname 参数是相对路径，比如
	// 		java.lang.Object -> java/lang/Object.class
	readClass(classname string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, constants.PathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, constants.WildCard) {
		return newWildCardEntry(path)
	}

	lowerCasePath := strings.ToLower(path)
	if strings.HasSuffix(lowerCasePath, constants.SuffixJar) ||
		strings.HasSuffix(lowerCasePath, constants.SuffixZip) {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
