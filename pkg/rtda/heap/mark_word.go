package heap

import (
	"sync"
)

var hashCodeGenerator int32 = 0

type MarkWord struct {
	hashCode int32
	mutex    sync.Mutex
}
