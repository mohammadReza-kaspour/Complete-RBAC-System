package domainerr

type Error struct {
	code    int
	message string
	cause   error
}

func New(code int, message string) *Error {
	return &Error{
		code:    code,
		message: message,
	}
}

func NewWithCause(code int, message string, cause error) *Error {
	return &Error{
		code:    code,
		message: message,
		cause:   cause,
	}
}

func (e *Error) Error() string {
	return e.message
}
func (e *Error) Code() int {
	return e.code
}
func (e *Error) Message() string {
	return e.message
}
func (e *Error) Cause() error {
	return e.cause
}

// it works on all errors not just custom error and get code if it has
func Code(err error) (int, bool) {
	e, ok := err.(*Error)
	if !ok {
		return 0, false
	}

	return e.code, true
}

// it works on all domain errors that say if error has current code or not
func Has(err error, code int) bool {
	c, ok := Code(err)
	return ok && c == code
}
