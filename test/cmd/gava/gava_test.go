package gava

import (
	"os"
	"os/exec"
	"strings"
	"testing"

	"jvm/pkg/constants"
	"jvm/test"
)

func TestGavaVersion(t *testing.T) {
	cmd := exec.Command(constants.Gava, "-version")

	test.CommonTestCommand(t, cmd, func(output string) {
		if !strings.HasPrefix(output, constants.Version) {
			t.Error("版本信息获取失败")
		}
	})
}

func TestGavaHelp(t *testing.T) {
	cmd := exec.Command(constants.Gava, "-help")

	test.CommonTestCommand(t, cmd, func(output string) {
		if !strings.HasPrefix(output, "Usage") {
			t.Error("获取帮助信息失败")
		}
	})
}

func TestStartJVM(t *testing.T) {
	cmd := exec.Command(constants.Gava, "-cp", os.Getenv("CLASSPATH"), "java.lang.String", "arg1", "arg2")

	test.CommonTestCommand(t, cmd, func(output string) {
		if !strings.Contains(output, "java.lang.String") ||
			!strings.Contains(output, "JavaVirtualMachines") {
			t.Error("启动虚拟机失败")
		}
	})
}
