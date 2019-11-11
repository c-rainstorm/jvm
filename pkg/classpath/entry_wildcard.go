package classpath

import (
	"os"
	"path/filepath"
	"strings"

	"jvm/pkg/constants"
)

func newWildCardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]

	var compositeEntry []Entry

	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// JVM 的通配符只支持扫到指定目录里的Jar文件，不支持递归搜索，不包含目录中的 class 文件
		// https://en.wikipedia.org/wiki/Classpath_(Java)#Adding_all_JAR_files_in_a_directory
		//      In Java 6 and higher, one can add 【all jar-files in a specific directory】 to the classpath using wildcard notation.
		// https://docs.oracle.com/javase/8/docs/technotes/tools/windows/classpath.html#A1100762
		// https://stackoverflow.com/questions/219585/including-all-the-jars-in-a-directory-within-the-java-classpath
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		lowerCasePath := strings.ToLower(path)
		if strings.HasSuffix(lowerCasePath, constants.SuffixJar) {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)
		}

		return nil
	}

	filepath.Walk(baseDir, walkFn)

	return compositeEntry
}
