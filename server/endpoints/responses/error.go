package responses

type Error struct {
	Error string `json:"error"`
}

// Builder Object for Error
type ErrorBuilder struct {
	err error
}

// Constructor for ErrorBuilder
func NewErrorBuilder() *ErrorBuilder {
	o := new(ErrorBuilder)
	return o
}

// Build Method which creates Error
func (b *ErrorBuilder) Build() *Error {
	o := new(Error)
	o.Error = b.err.Error()
	return o
}

// Setter method for the field err of type error in the object ErrorBuilder
func (e *ErrorBuilder) SetError(err error) {
	e.err = err
}
