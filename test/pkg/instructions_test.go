package pkg

import (
	"os"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
	"jvm/pkg/instructions"
)

func TestIOTA(t *testing.T) {
	t.Log(instructions.OpcIConst5)
	t.Log(instructions.OpcFLoad3)
	t.Log(instructions.OpcSALoad)
	t.Log(instructions.OpcSAStore)
	t.Log(instructions.OpcBreakPoint)
	t.Log(instructions.OpcImpDep1)
	t.Logf("当前指令未实现：%X", uint8(253))

}

func TestInterpret(t *testing.T) {
	// todo 测试没过

	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.GaussTest",
	}

	gava.Main()
}
