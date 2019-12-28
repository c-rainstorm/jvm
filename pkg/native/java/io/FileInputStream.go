package io

import (
	"jvm/pkg/native"
	"jvm/pkg/rtda"
)

func init() {
	native.MethodRegistry.Registry("java/io/FileInputStream", "initIDs", "()V", initIDs)
}

func initIDs(frame *rtda.Frame) {

}
