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
		interpret(mainMethod, cmd.args)
	} else {
		log.Print("main method not found!")
	}
}

func interpret(method *heap.Method, args []string) {
	thread := rtda.NewThread()
	frame := thread.NewFrame(method)

	// main 方法是 static 所以直接放第一个位置
	frame.LocalVars().SetRef(0, createArgsArray(method.Class().ClassLoader(), args))

	thread.PushFrame(frame)

	startInterpret(thread)
}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.Object {
	length := len(args)
	strArr := loader.LoadClass("java/lang/String").ArrayClass().NewArray(int32(length))
	for i, arg := range args {
		strArr.Set(int32(i), heap.JString(loader, arg))
	}
	return strArr
}

func catchError(thread *rtda.Thread) {
	if r := recover(); r != nil {
		for !thread.IsStackEmpty() {
			frame := thread.PopFrame()
			method := frame.Method()
			class := method.Class()
			log.Errorf(">> pc:%4d %v.%v%v \n", frame.NextPC(), class.Name(), method.Name(), method.Descriptor())
		}
		panic(r)
	}
}

func startInterpret(thread *rtda.Thread) {
	reader := &base.ByteCodeReader{}

	defer catchError(thread)
	for {
		frame := thread.CurrentFrame()
		pc := frame.NextPC()
		thread.SetPC(pc)

		reader.Reset(frame.Method().Code(), pc)
		opCode := reader.ReadUint8()
		inst := instructions.New(opCode)
		if inst == nil {
			log.Printf("inst not found, 0x%X", opCode)
			break
		}

		if global.Verbose {
			method := frame.Method()
			log.Infof("%v.%v() #%2d %T %v", method.Class().Name(), method.Name(), pc, inst, inst)
		}

		inst.FetchOperands(reader)
		// 读完指令以后，下次执行的PC会向后移动
		frame.SetNextPC(reader.PC())

		inst.Execute(frame)

		if thread.IsStackEmpty() {
			break
		}
	}
}
