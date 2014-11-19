package rainmaker

type ResponseReadError struct {
	internalError error
}

func NewResponseReadError(err error) ResponseReadError {
	return ResponseReadError{
		internalError: err,
	}
}

func (err ResponseReadError) Error() string {
	return "Rainmaker ResponseReadError: " + err.internalError.Error()
}

type ResponseBodyUnmarshalError struct {
	internalError error
}

func NewResponseBodyUnmarshalError(err error) ResponseBodyUnmarshalError {
	return ResponseBodyUnmarshalError{
		internalError: err,
	}
}

func (err ResponseBodyUnmarshalError) Error() string {
	return "Rainmaker ResponseBodyUnmarshalError: " + err.internalError.Error()
}

type RequestBodyMarshalError struct {
	internalError error
}

func NewRequestBodyMarshalError(err error) RequestBodyMarshalError {
	return RequestBodyMarshalError{internalError: err}
}

func (err RequestBodyMarshalError) Error() string {
	return "Rainmaker RequestBodyMarshalError: " + err.internalError.Error()
}

type RequestConfigurationError struct {
	internalError error
}

func NewRequestConfigurationError(err error) RequestConfigurationError {
	return RequestConfigurationError{internalError: err}
}

func (err RequestConfigurationError) Error() string {
	return "Rainmaker RequestConfigurationError: " + err.internalError.Error()
}

type RequestHTTPError struct {
	internalError error
}

func NewRequestHTTPError(err error) RequestHTTPError {
	return RequestHTTPError{internalError: err}
}

func (err RequestHTTPError) Error() string {
	return "Rainmaker RequestHTTPError: " + err.internalError.Error()
}
