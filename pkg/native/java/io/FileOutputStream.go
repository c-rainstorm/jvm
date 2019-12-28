package io

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry("java/io/FileOutputStream", "initIDs", "()V", FOS_initIDs)
}

func FOS_initIDs(frame *rtda.Frame) {

}
