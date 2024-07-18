package tango_errors

import "fmt"

type ModelError struct {
	ModelName string
	Message   string
	Code      int
}

func (merr *ModelError) Error() string {
	return fmt.Sprintf("model %s: (CODE: %d) err %s", merr.ModelName, merr.Code, merr.Message)
}
