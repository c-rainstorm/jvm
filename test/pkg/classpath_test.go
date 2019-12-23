package pkg

import (
	"os"
	"path/filepath"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
)

func TestDirEntry(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-Xjre", filepath.Join(os.Getenv(global.JavaHome), "jre"),
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.HelloWorld", "arg1", "arg2",
	}

	gava.Main()
}

func TestJarEntry(t *testing.T) {
	os.Args = []string{
		global.Gava,
		"java.lang.String", "arg1", "arg2",
	}

	gava.Main()
}
