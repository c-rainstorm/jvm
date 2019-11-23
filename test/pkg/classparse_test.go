package pkg

import (
	"os"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
)

func TestStringParse(t *testing.T) {
	os.Args = []string{
		global.Gava,
		"-v",
		"java.lang.String",
	}

	gava.Main()
}

func TestClassParse(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.ClassFileParseTest", "arg1", "arg2",
	}

	gava.Main()
}

func TestInterface1Parse(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.Interface1", "arg1", "arg2",
	}

	gava.Main()
}

func TestInterface2Parse(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.Interface2", "arg1", "arg2",
	}

	gava.Main()
}

func TestAnnoParse(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.Anno", "arg1", "arg2",
	}

	gava.Main()
}
