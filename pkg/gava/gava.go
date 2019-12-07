package gava

import (
	"strings"

	"jvm/pkg/classpath"
	"jvm/pkg/global"
	"jvm/pkg/instructions"
	"jvm/pkg/instructions/base"
	"jvm/pkg/logger"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
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

	classLoader := heap.NewClassLoader(cp)
	classname := strings.Replace(cmd.class, global.Dot, global.Slash, -1)
	if global.Verbose {
		log.Infof("classname: %s", classname)
	}

	mainClass := classLoader.LoadClass(classname)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		log.Print("main method not found!")
	}
}

func interpret(method *heap.Method) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)
	thread.PushFrame(frame)

	startInterpret(thread, method.Code())
}

func catchError(frame *rtda.Frame) {
	if r := recover(); r != nil {
		log.Printf("LocalVars:%v\n", frame.LocalVars())
		log.Printf("OperandStack: %v", frame.OperandStack())
		panic(r)
	}
}

func startInterpret(thread *rtda.Thread, byteCode []byte) {
	frame := thread.PopFrame()
	reader := &base.ByteCodeReader{}

	defer catchError(frame)
	for {
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(byteCode, pc)
		opCode := reader.ReadUint8()
		inst := instructions.New(opCode)
		inst.FetchOperands(reader)
		// 读完指令以后，下次执行的PC会向后移动
		frame.SetNextPC(reader.PC())

		log.Printf("pc:%v inst:%T %v\n", pc, inst, inst)
		inst.Execute(frame)
	}
}
