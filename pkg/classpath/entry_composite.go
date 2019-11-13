package classpath

import (
	"strings"

	"jvm/pkg/exception"
	"jvm/pkg/global"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry []Entry
	for _, path := range strings.Split(pathList, global.PathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry)
	}
	return compositeEntry
}

func (self CompositeEntry) readClass(classname string) ([]byte, Entry, error) {
	for _, entry := range self {
		dataBytes, from, err := entry.readClass(classname)
		if err == nil {
			return dataBytes, from, nil
		}
	}

	return nil, nil, exception.ClassNotFound(classname)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))

	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, global.PathListSeparator)
}
