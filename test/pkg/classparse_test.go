package pkg

import (
	"os"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
)

func TestMagicParse(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.ClassFileParseTest", "arg1", "arg2",
	}

	gava.Main()
}