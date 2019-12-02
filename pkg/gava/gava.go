package gava

import (
	"strings"

	"jvm/pkg/classfile"
	"jvm/pkg/classpath"
	"jvm/pkg/global"
	"jvm/pkg/instructions"
	"jvm/pkg/instructions/base"
	"jvm/pkg/logger"
	"jvm/pkg/rtda"
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

	mainMethod := findMainMethod(classFile)

	if mainMethod != nil {
		interpret(mainMethod)
	} else {
		log.Print("main method not found!")
	}
}

func interpret(method *classfile.MethodMemberInfo) {
	codeAttr := method.CodeAttr()
	maxLocals := codeAttr.MaxLocals()
	maxStack := codeAttr.MaxStack()
	code := codeAttr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(uint(maxLocals), uint(maxStack))
	thread.PushFrame(frame)

	startInterpret(thread, code)
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

func findMainMethod(classFile *classfile.ClassFile) *classfile.MethodMemberInfo {
	for _, m := range *classFile.Methods() {
		if m.Name() == "main" && m.Descriptor() == "([Ljava/lang/String;)V" {
			return m
		}
	}
	return nil
}
