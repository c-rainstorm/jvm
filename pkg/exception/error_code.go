package exception

type ErrorCode int

const (
	ErrClassNotFound ErrorCode = iota
)

func (this ErrorCode) String() string {
	switch this {
	case ErrClassNotFound:
		return "ClassNotFound"
	default:
		return string(this)
	}
}
