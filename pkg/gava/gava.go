package gava

import (
	"jvm/pkg/global"
	"jvm/pkg/logger"
)

var log = logger.NewLogrusLogger()

func Main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		log.Info(global.Version)
	} else if cmd.helpFlag || cmd.class == global.EmptyString {
		printUsage()
	} else {
		newJVM(cmd).start()
	}
}
