package gava

import (
	"strings"

	"jvm/pkg/classfile"
	"jvm/pkg/classpath"
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
		startJVM(cmd)
	}
}

func startJVM(cmd *Cmd) {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	if global.Verbose {
		log.Info(cmd)
	}

	classname := strings.Replace(cmd.class, global.Dot, global.Slash, -1)
	if global.Verbose {
		log.Infof("classname: %s", classname)
	}

	classData, _, err := cp.ReadClass(classname)
	if err != nil {
		log.Errorln("Could not find or load main class " + cmd.class)
		return
	}

	classFile := classfile.Parse(classData)

	if global.Verbose {
		log.Infof("class data: %s", classFile)
	}
}
