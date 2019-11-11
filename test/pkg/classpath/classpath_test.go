package classpath

import (
	"jvm/pkg/constants"
	"jvm/pkg/gava"
	"os"
	"path/filepath"
	"testing"
)

func TestDirEntry(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		constants.Gava,
		"-Xjre", filepath.Join(os.Getenv(constants.JavaHome), "jre"),
		"-cp", cp,
		"me.rainstorm.jvm.HelloWorld", "arg1", "arg2",
	}

	gava.Main()
}

func TestJarEntry(t *testing.T) {
	os.Args = []string{
		constants.Gava,
		"java.lang.String", "arg1", "arg2",
	}

	gava.Main()
}
