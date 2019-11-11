package exception

import (
	"fmt"
)

// 通用异常格式
// 第一个参数是异常码，定义在 pkg/exception/error_code.go:3 中
// 第二个参数是错误消息，定义在
const errorMsgPattern string = "[%s] - %s"

func ClassNotFound(classname string) error {
	return Error(ErrClassNotFound, classname)
}

func Error(code ErrorCode, msg string) error {
	return fmt.Errorf(errorMsgPattern, code, msg)
}
