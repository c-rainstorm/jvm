package pkg

import (
	"os"
	"sync"
	"testing"
	"time"

	"jvm/pkg/gava"
	"jvm/pkg/global"
	"jvm/pkg/rtda/heap"
)

func TestHelloWorld(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.HelloWorld", "Hello world!",
	}

	gava.Main()
}

func TestNative(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		// "-v",
		"me.rainstorm.jvm.NativeTest",
	}

	gava.Main()
}

func TestStringAppend(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		"-v",
		"me.rainstorm.jvm.StringAppend",
	}

	gava.Main()
}

var wait sync.WaitGroup

func TestHashCode(t *testing.T) {
	var size = 100
	wait.Add(size)
	for i := 0; i < size; i++ {
		go printHashCode(t)
	}
	time.Sleep(10 * time.Second)
}

func printHashCode(t *testing.T) {
	wait.Done()
	wait.Wait()
	obj := &heap.BaseObject{}
	hc1 := obj.HashCode()
	hc2 := obj.HashCode()
	if hc1 != hc2 {
		t.Error("hashCode not equal: ", hc1, ", ", hc2)
	} else {
		t.Log(hc1, ", ", hc2)
	}
}

func TestBoxTest(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		// "-v",
		"me.rainstorm.jvm.BoxTest",
	}

	gava.Main()
}

func TestExceptionTest(t *testing.T) {
	cp := "/Users/chen/workspace/go/src/jvm/test/data/class"
	os.Args = []string{
		global.Gava,
		"-cp", cp,
		// "-v",
		"me.rainstorm.jvm.ExceptionThrowTest", "a",
	}

	gava.Main()
}
