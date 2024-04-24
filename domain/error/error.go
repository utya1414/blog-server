package error

type Error struct {
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

var (
	ErrUserAlreadyExists = &Error{Message: "ユーザーがすでに存在しています"}
)