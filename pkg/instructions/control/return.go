package control

import (
	"jvm/pkg/instructions/base"
	"jvm/pkg/rtda"
)

type Return struct {
	base.NoOperandsInstruction
}

func (this *Return) Execute(frame *rtda.Frame) {
	frame.Thread().PopFrame()
}

type AReturn struct {
	base.NoOperandsInstruction
}

func (this *AReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushRef(currentFrame.OperandStack().PopRef())
}

type IReturn struct {
	base.NoOperandsInstruction
}

func (this *IReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushInt(currentFrame.OperandStack().PopInt())
}

type LReturn struct {
	base.NoOperandsInstruction
}

func (this *LReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushLong(currentFrame.OperandStack().PopLong())
}

type FReturn struct {
	base.NoOperandsInstruction
}

func (this *FReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushFloat(currentFrame.OperandStack().PopFloat())
}

type DReturn struct {
	base.NoOperandsInstruction
}

func (this *DReturn) Execute(frame *rtda.Frame) {
	thread := frame.Thread()
	currentFrame := thread.PopFrame()
	invokerFrame := thread.TopFrame()
	invokerFrame.OperandStack().PushDouble(currentFrame.OperandStack().PopDouble())
}
