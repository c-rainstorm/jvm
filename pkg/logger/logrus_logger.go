package logger

import (
	"os"

	"github.com/sirupsen/logrus"
	"jvm/pkg/logger/formatter"
)

func NewLogrusLogger() *logrus.Logger {
	var log = logrus.New()

	// 设置日志格式为json格式
	log.SetFormatter(&formatter.DefaultLogrusFormatter)

	// 设置将日志输出到标准输出（默认的输出为stderr,标准错误）
	// 日志消息输出可以是任意的io.writer类型
	log.SetOutput(os.Stdout)

	// 设置日志级别为warn以上
	log.SetLevel(logrus.DebugLevel)

	// 打印方法名
	log.SetReportCaller(true)

	return log
}
