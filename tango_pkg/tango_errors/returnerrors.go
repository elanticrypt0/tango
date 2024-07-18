package tango_errors

// return Error as struct
func ReturnDefault(name, message string, code int) *DefaultError {
	return &DefaultError{
		Name:    name,
		Code:    code,
		Message: message,
	}
}

// return Error as struct
func ReturnModel(name, message string, code int) *ModelError {
	return &ModelError{
		ModelName: name,
		Code:      code,
		Message:   message,
	}
}
