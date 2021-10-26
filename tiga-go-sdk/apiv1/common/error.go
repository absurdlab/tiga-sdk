package commonv1

// WrapErrorResponse wraps ErrorResponse in an error implementation.
func WrapErrorResponse(resp *ErrorResponse) error {
	return ErrorResponseWrapper{resp: resp}
}

type ErrorResponseWrapper struct {
	resp *ErrorResponse
}

func (e ErrorResponseWrapper) Error() string {
	return e.resp.Error
}

func (e ErrorResponseWrapper) Response() *ErrorResponse {
	return e.resp
}
