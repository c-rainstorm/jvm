package test

import (
	"fmt"
	"os/exec"
	"testing"
)

func CommonTestCommand(t *testing.T, cmd *exec.Cmd, outputCheck func(string)) {
	out, err := cmd.CombinedOutput()
	output := string(out)

	if err != nil {
		t.Errorf("%v", err)
	}

	fmt.Println(output)

	outputCheck(output)
}