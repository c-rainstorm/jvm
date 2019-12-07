package heap

import "jvm/pkg/logger"

var log = logger.NewLogrusLogger()

type Object struct {
	class  *Class
	fields *Slots
}
