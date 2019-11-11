package log

import (
	"github.com/sirupsen/logrus"
	"jvm/pkg/exception"
	"jvm/pkg/logger"
	"testing"
)

var log = logger.NewLogrusLogger()

func TestFields(t *testing.T) {
	log.WithFields(logrus.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}

func TestError(t *testing.T) {
	log.Error(exception.ClassNotFound("java.lang.Object"))
}
