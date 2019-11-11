package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &DirEntry{absDir: absDir}
}

func (this *DirEntry) readClass(className string) ([]byte, Entry, error) {
	filename := filepath.Join(this.absDir, className)
	log.Infof("class absDir: %s", filename)

	data, err := ioutil.ReadFile(filename)
	return data, this, err
}

func (this *DirEntry) String() string {
	return this.absDir
}
