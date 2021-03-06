package gava

import (
	"strings"

	"jvm/pkg/classpath"
	"jvm/pkg/global"
	"jvm/pkg/instructions"
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

type JVM struct {
	loader     *heap.ClassLoader
	cmd        *Cmd
	mainThread *rtda.Thread
}

func newJVM(cmd *Cmd) *JVM {
	cp := classpath.Parse(cmd.XjreOption, cmd.cpOption)
	if global.Verbose {
		log.Info(cmd)
	}

	classLoader := heap.NewClassLoader(cp)

	return &JVM{
		loader:     classLoader,
		cmd:        cmd,
		mainThread: rtda.NewThread(),
	}
}

func (this *JVM) start() {
	// this.init()
	this.main()
}

func (this *JVM) init() {
	systemClass := this.loader.LoadClass("sun/misc/VM")

	base.InitClass(this.mainThread, systemClass)

	StartInterpret(this.mainThread)
}

func (this *JVM) main() {
	classname := strings.Replace(this.cmd.class, global.Dot, global.Slash, -1)
	if global.Verbose {
		log.Infof("classname: %s", classname)
	}

	mainClass := this.loader.LoadClass(classname)
	mainMethod := mainClass.GetMainMethod()

	if mainMethod != nil {
		frame := this.mainThread.NewFrame(mainMethod)

		// main 方法是 static 所以直接放第一个位置
		frame.LocalVars().SetRef(0, createArgsArray(this.loader, this.cmd.args))

		this.mainThread.PushFrame(frame)
		StartInterpret(this.mainThread)
	} else {
		log.Print("main method not found!")
	}
}

func createArgsArray(loader *heap.ClassLoader, args []string) *heap.ArrayObject {
	length := len(args)
	strArr := loader.LoadClass(global.JavaLangString).ArrayClass().NewArray(int32(length))
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

func StartInterpret(thread *rtda.Thread) {
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
