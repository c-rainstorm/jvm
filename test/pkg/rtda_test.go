package pkg

import (
	"os"
	"testing"

	"jvm/pkg/gava"
	"jvm/pkg/global"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
)

func TestLocalVars(t *testing.T) {
	localVars := rtda.NewLocalVars(100)

	intVal := int32(100)
	localVars.SetInt(0, intVal)
	intResult := localVars.GetInt(0)
	t.Logf("int: %v", intResult)
	if intVal != intResult {
		t.Errorf("target: %v, actual: %v", intVal, intResult)
	}

	floatVal := float32(3.14)
	localVars.SetFloat(1, floatVal)
	floatResult := localVars.GetFloat(1)
	t.Logf("float: %v", floatResult)
	if floatVal != floatResult {
		t.Errorf("target: %v, actual: %v", floatVal, floatResult)
	}

	longVal := int64(100000000000000)
	localVars.SetLong(2, longVal)
	longResult := localVars.GetLong(2)
	t.Logf("long: %v", longResult)
	if longVal != longResult {
		t.Errorf("target: %v, actual: %v", longVal, longResult)
	}

	doubleVal := float64(3.1445678765456783456)
	localVars.SetDouble(4, doubleVal)
	doubleResult := localVars.GetDouble(4)
	t.Logf("double: %v", doubleResult)
	if doubleVal != doubleResult {
		t.Errorf("target: %v, actual: %v", doubleVal, doubleResult)
	}

	refVal := &heap.Object{}
	localVars.SetRef(6, refVal)
	refResult := localVars.GetRef(6)
	t.Logf("ref: %v", refResult)
	if refVal != refResult {
		t.Errorf("target: %v, actual: %v", refVal, refResult)
	}
}

func TestOperandStack(t *testing.T) {
	operandStack := rtda.NewOperandStack(100)

	intVal := int32(100)
	operandStack.PushInt(intVal)
	intResult := operandStack.PopInt()
	t.Logf("int: %v", intResult)
	if intVal != intResult {
		t.Errorf("target: %v, actual: %v", intVal, intResult)
	}

	floatVal := float32(3.14)
	operandStack.PushFloat(floatVal)
	floatResult := operandStack.PopFloat()
	t.Logf("float: %v", floatResult)
	if floatVal != floatResult {
		t.Errorf("target: %v, actual: %v", floatVal, floatResult)
	}

	longVal := int64(123432343234)
	operandStack.PushLong(longVal)
	longResult := operandStack.PopLong()
	t.Logf("long: %v", longResult)
	if longVal != longResult {
		t.Errorf("target: %v, actual: %v", longVal, longResult)
	}

	doubleVal := float64(3.1445678765456783456)
	operandStack.PushDouble(doubleVal)
	doubleResult := operandStack.PopDouble()
	t.Logf("double: %v", doubleResult)
	if doubleVal != doubleResult {
		t.Errorf("target: %v, actual: %v", doubleVal, doubleResult)
	}

	refVal := &heap.Object{}
	operandStack.PushRef(refVal)
	refResult := operandStack.PopRef()
	t.Logf("ref: %v", refResult)
	if refVal != refResult {
		t.Errorf("target: %v, actual: %v", refVal, refResult)
	}
}

func TestClassAndField(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.MyObject",
	}

	gava.Main()
}

func TestMethodInvoke(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.InvokeTest",
	}

	gava.Main()
}

func TestFibonacci(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.FibonacciTest",
	}

	gava.Main()
}

func TestClinitTest(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.ClinitTest",
	}

	gava.Main()
}

func TestInt64ToInt32(t *testing.T) {
	longVal := int64(123432343234)
	low32 := int32(uint32(longVal))
	high32 := int32(uint32(longVal >> 32))
	t.Logf("target: 0x%X, 0x%X, 0x%X", longVal, low32, high32)
	t.Logf("target: 0x%X", int64(uint32(low32))|(int64(uint32(high32))<<32))
}
