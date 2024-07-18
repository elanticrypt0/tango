package tango_errors

import "fmt"

type DefaultError struct {
	Name    string
	Message string
	Code    int
}

func (merr *DefaultError) Error() string {
	return fmt.Sprintf("model %s: (CODE: %d) err %s", merr.Name, merr.Code, merr.Message)
}
