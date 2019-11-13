package gava

import (
	"flag"
	"fmt"
	"os"

	"jvm/pkg/global"
)

type Cmd struct {
	verboseFlag bool
	helpFlag    bool
	versionFlag bool
	cpOption    string
	XjreOption  string
	class       string
	args        []string
}

func parseCmd() *Cmd {
	cmd := &Cmd{}

	flag.Usage = printUsage
	flag.BoolVar(&cmd.verboseFlag, "verbose", false, "print detail message")
	flag.BoolVar(&cmd.verboseFlag, "v", false, "print detail message")
	flag.BoolVar(&cmd.helpFlag, "help", false, "print help message")
	flag.BoolVar(&cmd.helpFlag, "?", false, "print help message")
	flag.BoolVar(&cmd.versionFlag, "version", false, "print version and exit")
	flag.StringVar(&cmd.cpOption, "classpath", "", "classpath")
	flag.StringVar(&cmd.cpOption, "cp", "", "classpath")
	flag.StringVar(&cmd.XjreOption, "Xjre", "", "path to jre")
	flag.Parse()

	args := flag.Args()
	if len(args) > 0 {
		cmd.class = args[0]
		cmd.args = args[1:]
	}

	global.Verbose = cmd.verboseFlag

	return cmd
}

func printUsage() {
	log.Infof("Usage: %s [-options] class [args...]\n", os.Args[0])
}

func (this *Cmd) String() string {
	return fmt.Sprintf("Cmd{verboseFlag: %v, helpFlag: %v, versionFlag: %v, cpOption: %s, XjreOption: %s, class: %s, args: %s}",
		this.verboseFlag, this.helpFlag, this.versionFlag, this.cpOption, this.XjreOption, this.class, this.args)
}
