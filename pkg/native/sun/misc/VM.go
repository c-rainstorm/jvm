package misc

import (
	"os"

	"jvm/pkg/global"
	"jvm/pkg/native"
	"jvm/pkg/rtda"
	"jvm/pkg/rtda/heap"
	"jvm/pkg/rtda/invoke"
)

func init() {
	prepareProperties()
	native.MethodRegistry.Registry("sun/misc/VM", "initialize", "()V", initialize)
}

func initialize(frame *rtda.Frame) {
	// systemClass := frame.Method().Class().ClassLoader().LoadClass("java/lang/System")
	// //  private static void initializeSystemClass();
	// initializeSystemClassMethod := systemClass.GetMethod("initializeSystemClass", "()V")
	//
	// invoke.InvokeMethod(frame, initializeSystemClassMethod)
	// 添加一个元素到 savedProps
	VM := frame.Method().Class()
	savedProps := VM.GetField(nil, "savedProps", "Ljava/util/Properties;").(*heap.NormalObject)
	for key, value := range Properties {
		invoke.SetProperty(frame.Thread(), savedProps, key, value)
	}
}

var Properties = make(map[string]string)

func prepareProperties() {
	Properties["java.runtime.name"] = "Java(TM) SE Runtime Environment"
	Properties["sun.boot.library.path"] = "/Library/Java/JavaVirtualMachines/jdk..."
	Properties["java.vm.version"] = global.Version
	Properties["path.separator"] = string(os.PathListSeparator)
	Properties["file.separator"] = string(os.PathSeparator)
	// java.vm.version=25.172-b11
	// gopherProxySet=false
	// java.vm.vendor=Oracle Corporation
	// java.vendor.url=http://java.oracle.com/

	// java.vm.name=Java HotSpot(TM) 64-Bit Server VM
	// file.encoding.pkg=sun.io
	// user.country=CN
	// sun.java.launcher=SUN_STANDARD
	// sun.os.patch.level=unknown
	// java.vm.specification.name=Java Virtual Machine Specification
	// user.dir=/Users/chen/workspace/git/furry-octo-...
	// java.runtime.version=1.8.0_172-b11
	// java.awt.graphicsenv=sun.awt.CGraphicsEnvironment
	// java.endorsed.dirs=/Library/Java/JavaVirtualMachines/jdk...
	// os.arch=x86_64
	// visualvm.id=574343783456060
	// java.io.tmpdir=/var/folders/s7/mpzn0l2d3k52rh0zlm4md...
	// line.separator=
	//
	// 	java.vm.specification.vendor=Oracle Corporation
	// os.name=Mac OS X
	// sun.jnu.encoding=UTF-8
	// java.library.path=/Users/chen/Library/Java/Extensions:/...
	// java.specification.name=Java Platform API Specification
	// java.class.version=52.0
	// sun.management.compiler=HotSpot 64-Bit Tiered Compilers
	// os.version=10.13.6
	// user.home=/Users/chen
	// user.timezone=
	// 	java.awt.printerjob=sun.lwawt.macosx.CPrinterJob
	// file.encoding=UTF-8
	// java.specification.version=1.8
	// user.name=chen
	// java.class.path=/Library/Java/JavaVirtualMachines/jdk...
	// java.vm.specification.version=1.8
	// sun.arch.data.model=64
	// java.home=/Library/Java/JavaVirtualMachines/jdk...
	// sun.java.command=me.rainstorm.fod.test.ExceptionThrowT...
	// java.specification.vendor=Oracle Corporation
	// user.language=en
	// awt.toolkit=sun.lwawt.macosx.LWCToolkit
	// java.vm.info=mixed mode
	// java.version=1.8.0_172
	// java.ext.dirs=/Users/chen/Library/Java/Extensions:/...
	// sun.boot.class.path=/Library/Java/JavaVirtualMachines/jdk...
	// java.vendor=Oracle Corporation
	// file.separator=/
	// java.vendor.url.bug=http://bugreport.sun.com/bugreport/
	// sun.cpu.endian=little
	// sun.io.unicode.encoding=UnicodeBig
	// sun.cpu.isalist=
}
