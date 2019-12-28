package io

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry("java/io/FileDescriptor", "initIDs", "()V", FD_initIDs)
}

func FD_initIDs(frame *rtda.Frame) {

}
