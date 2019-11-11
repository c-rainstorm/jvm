package gava

import (
	"jvm/pkg/classpath"
	"jvm/pkg/constants"
	"jvm/pkg/logger"
	"strings"
)

var log = logger.NewLogrusLogger()

func Main() {
	cmd := parseCmd()

	if cmd.versionFlag {
		log.Info(constants.Version)
	} else if cmd.helpFlag || cmd.class == constants.EmptyString {
		printUsage()
	} else {
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	log.Infof("classpath:%s cmd:%v", cp, cmd)

	classname := strings.Replace(cmd.class, constants.Dot, constants.Slash, -1)
	log.Infof("classname: %s", classname)

	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		log.Errorln("Could not find or load main class " + cmd.class)
		return
	}

	log.Infof("class data: %v", classData)
}
