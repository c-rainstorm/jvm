package exception

import (
	"fmt"
	"jvm/pkg/exception"
	"strings"
	"testing"
)

func TestError(t *testing.T) {
	err := exception.Error(exception.ErrClassNotFound, "java.lang.Object")

	fmt.Println(err)

	if err == nil {
		t.Error("ErrClassNotFound error 生成失败")
		return
	}

	errStr := err.Error()
	if !strings.HasPrefix(errStr, "[ClassNotFound] - java.lang.Object") {
		t.Error("异常信息有误")
	}
}

func TestClassNotFound(t *testing.T) {
	err := exception.ClassNotFound("java.lang.Object")

	fmt.Println(err)

	if err == nil {
		t.Error("ErrClassNotFound error 生成失败")
		return
	}

	errStr := err.Error()
	if !strings.HasPrefix(errStr, "[ClassNotFound] - java.lang.Object") {
		t.Error("异常信息有误")
	}
}
