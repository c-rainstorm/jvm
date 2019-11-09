package gava

import (
	"fmt"
	"jvm/pkg/constants"
	"os"
	"os/exec"
	"strings"
	"testing"
)

const GAVA string = "gava"

func TestGavaVersion(t *testing.T) {
	cmd := exec.Command(GAVA, "-version")

	testCommand(t, cmd, func(output string) {
		if !strings.HasPrefix(output, constants.Version) {
			t.Error("版本信息获取失败")
		}
	})
}

func TestGavaHelp(t *testing.T) {
	cmd := exec.Command(GAVA, "-help")

	testCommand(t, cmd, func(output string) {
		if !strings.HasPrefix(output, "Usage") {
			t.Error("获取帮助信息失败")
		}
	})
}

func TestStartJVM(t *testing.T) {
	cmd := exec.Command(GAVA, "-cp", os.Getenv("CLASSPATH"), "java.lang.String", "arg1", "arg2")

	testCommand(t, cmd, func(output string) {
		if !strings.Contains(output, "java.lang.String") ||
			!strings.Contains(output, "JavaVirtualMachines") {
			t.Error("启动虚拟机失败")
		}
	})
}

func testCommand(t *testing.T, cmd *exec.Cmd, outputCheck func(string)) {
	out, err := cmd.CombinedOutput()
	output := string(out)

	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Println(output)

	outputCheck(output)
}
