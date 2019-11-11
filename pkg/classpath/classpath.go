package classpath

import (
	"jvm/pkg/constants"
	"jvm/pkg/logger"
	"os"
	"path/filepath"
)

var log = logger.NewLogrusLogger()

type Classpath struct {
	bootClasspath Entry
	extClasspath  Entry
	userClasspath Entry
}

func (this *Classpath) parseBootAndExtClasspath(jreOption string) {
	jrePath := getJreDir(jreOption)

	// jre/lib/*
	jreLibPath := filepath.Join(jrePath, "lib", constants.WildCard)
	this.bootClasspath = newEntry(jreLibPath)

	// jre/lib/ext/*
	jreExtLibPath := filepath.Join(jrePath, "lib", "ext", constants.WildCard)
	this.extClasspath = newEntry(jreExtLibPath)
}

func getJreDir(jreOption string) string {
	if jreOption != constants.EmptyString && exists(jreOption) {
		return jreOption
	}

	if exists("./jre") {
		return "./jre"
	}

	if javaHome := os.Getenv(constants.JavaHome); javaHome != constants.EmptyString {
		return filepath.Join(javaHome, "jre")
	}

	panic("Can not find jre folder!")
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}

	return true
}

func (this *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == constants.EmptyString {
		cpOption = constants.Dot
	}

	this.userClasspath = newEntry(cpOption)
}

func (this *Classpath) ReadClass(classname string) ([]byte, Entry, error) {
	classname = classname + constants.SuffixClass

	if dataBytes, entry, err := this.bootClasspath.readClass(classname); err == nil {
		return dataBytes, entry, err
	}

	log.Info("bootClasspath not found")

	if dataBytes, entry, err := this.extClasspath.readClass(classname); err == nil {
		return dataBytes, entry, err
	}

	log.Info("extClasspath not found")

	return this.userClasspath.readClass(classname)
}

func (this *Classpath) String() string {
	return this.userClasspath.String()
}

func Parse(jreOption, cpOption string) *Classpath {
	cp := &Classpath{}
	cp.parseBootAndExtClasspath(jreOption)
	cp.parseUserClasspath(cpOption)
	return cp
}
