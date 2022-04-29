package bots

type ErrorWithStatusCode struct {
	StatusCode int
	Err        error
}

func (e *ErrorWithStatusCode) Error() string {
	return e.Err.Error()
}
