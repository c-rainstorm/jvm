package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"

	"jvm/pkg/exception"
)

type ZipEntry struct {
	absDir string
}

func newZipEntry(path string) *ZipEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}

	return &ZipEntry{absDir: absDir}
}

func (self *ZipEntry) readClass(classname string) ([]byte, Entry, error) {
	reader, err := zip.OpenReader(self.absDir)
	if err != nil {
		return nil, nil, nil
	}

	// 等同 Java 的 finally 块的作用
	defer reader.Close()

	// 遍历包中所有的文件
	for _, file := range reader.File {
		if classname == file.Name {
			readCloser, err := file.Open()
			if err != nil {
				return nil, nil, nil
			}

			defer readCloser.Close()
			dataBytes, err := ioutil.ReadAll(readCloser)
			if err != nil {
				return nil, nil, nil
			}

			return dataBytes, self, err
		}
	}

	return nil, nil, exception.ClassNotFound(classname)
}

func (self *ZipEntry) String() string {
	return self.absDir
}
