package custerror

type CustomError interface {
	error
	CustomCode() uint16
	HTTP() uint16
	CodeExpr() string
}

const (
	BadRequestErrorCode   = 400
	UnauthorizedErrorCode = 401
	ForbiddenErrorCode    = 403
	NotFoundErrorCode     = 404
	DataConflictErrorCode = 409
	ValidationErrorCode   = 422
	ServerErrorCode       = 500
)

var (
	_ CustomError = &GeneralError{}
	_ CustomError = &ServerError{}
	_ CustomError = &BadRequestError{}
	_ CustomError = &UnauthorizedError{}
	_ CustomError = &ForbiddenError{}
	_ CustomError = &NotFoundError{}
)

// ServerError defines common 500 status code error that occurred due to the server.
type GeneralError struct {
	httpCode uint16
	codeExpr string
	message  string
}

func NewGeneralError(httpCode uint16, codeExpr, message string) *GeneralError {
	return &GeneralError{httpCode: httpCode, codeExpr: codeExpr, message: message}
}

func (e *GeneralError) Error() string {
	return e.message
}

func (e *GeneralError) CustomCode() uint16 {
	return 0
}

func (e *GeneralError) CodeExpr() string {
	return e.codeExpr
}

func (e *GeneralError) HTTP() uint16 {
	return e.httpCode
}

// ServerError defines common 500 status code error.
type ServerError struct {
	desc string
}

func NewServerError(err error) *ServerError {
	return &ServerError{desc: err.Error()}
}

func (e *ServerError) Error() string {
	return e.desc
}

func (e *ServerError) CustomCode() uint16 {
	return 0
}

func (e *ServerError) CodeExpr() string {
	return "UNEXPECTED_ERROR"
}

func (e *ServerError) HTTP() uint16 {
	return ServerErrorCode
}

// BadRequestError defines 400 status code error.
type BadRequestError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewBadRequestError(code uint16, codeExpr, desc string) *BadRequestError {
	return &BadRequestError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *BadRequestError) Error() string {
	return e.desc
}

func (e *BadRequestError) CustomCode() uint16 {
	return e.code
}

func (e *BadRequestError) CodeExpr() string {
	return e.codeExpr
}

func (e *BadRequestError) HTTP() uint16 {
	return BadRequestErrorCode
}

// UnauthorizedError defines 401 status code error.
type UnauthorizedError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewUnauthorizedError(code uint16, codeExpr, desc string) *UnauthorizedError {
	return &UnauthorizedError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *UnauthorizedError) Error() string {
	return e.desc
}

func (e *UnauthorizedError) CustomCode() uint16 {
	return e.code
}

func (e *UnauthorizedError) CodeExpr() string {
	return e.codeExpr
}

func (e *UnauthorizedError) HTTP() uint16 {
	return UnauthorizedErrorCode
}

// ForbiddenError defines 403 status code error.
type ForbiddenError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewForbiddenError(code uint16, codeExpr, desc string) *ForbiddenError {
	return &ForbiddenError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *ForbiddenError) Error() string {
	return e.desc
}

func (e *ForbiddenError) CustomCode() uint16 {
	return e.code
}

func (e *ForbiddenError) CodeExpr() string {
	return e.codeExpr
}

func (e *ForbiddenError) HTTP() uint16 {
	return ForbiddenErrorCode
}

// NotFoundError defines 404 status code error.
type NotFoundError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewNotFoundError(code uint16, codeExpr, desc string) *NotFoundError {
	return &NotFoundError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *NotFoundError) Error() string {
	return e.desc
}

func (e *NotFoundError) CustomCode() uint16 {
	return e.code
}

func (e *NotFoundError) CodeExpr() string {
	return e.codeExpr
}

func (e *NotFoundError) HTTP() uint16 {
	return NotFoundErrorCode
}

// DataConflictError defines 409 status code error.
type DataConflictError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewDataConflictError(code uint16, codeExpr, desc string) *DataConflictError {
	return &DataConflictError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *DataConflictError) Error() string {
	return e.desc
}

func (e *DataConflictError) CustomCode() uint16 {
	return e.code
}

func (e *DataConflictError) CodeExpr() string {
	return e.codeExpr
}

func (e *DataConflictError) HTTP() uint16 {
	return DataConflictErrorCode
}

// DataConflictError defines 409 status code error.
type ValidationError struct {
	code     uint16
	codeExpr string
	desc     string
}

func NewValidationError(code uint16, codeExpr, desc string) *ValidationError {
	return &ValidationError{code: code, codeExpr: codeExpr, desc: desc}
}

func (e *ValidationError) Error() string {
	return e.desc
}

func (e *ValidationError) CustomCode() uint16 {
	return e.code
}

func (e *ValidationError) CodeExpr() string {
	return e.codeExpr
}

func (e *ValidationError) HTTP() uint16 {
	return ValidationErrorCode
}
