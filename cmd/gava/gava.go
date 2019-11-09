package main

import (
	"fmt"
	"jvm/pkg/constants"
)

func main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		fmt.Println(constants.Version)
	} else if cmd.helpFlag || cmd.class == "" {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	fmt.Printf("classpath:%s class:%s args:%v\n",
		cmd.cpOption, cmd.class, cmd.args)
}
