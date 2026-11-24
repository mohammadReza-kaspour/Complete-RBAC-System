package ResultPattern

type ResultBase[T any] struct {
	IsSuccess  bool
	TraceId    string
	StatusCode int
	ErrorCode  int
	Message    string
	Data       T
}

type ErrorBase[T any] struct {
	ResultBase[T]
}

func (e *ErrorBase[T]) Error() string {
	return e.Message
}
