package pkg

import (
	"os"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
)

func TestIOTA(t *testing.T) {
	t.Log(global.OpcIConst5)
	t.Log(global.OpcFLoad3)
	t.Log(global.OpcSALoad)
	t.Log(global.OpcSAStore)
	t.Log(global.OpcBreakPoint)
	t.Log(global.OpcImpDep1)
	t.Logf("当前指令未实现：%X", uint8(253))

}

func TestGauss(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.GaussTest",
	}

	gava.Main()
}
