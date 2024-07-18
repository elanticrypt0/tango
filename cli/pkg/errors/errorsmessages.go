package errors

import (
	"fmt"
)

func FileNotExistError(filepath string) string {
	return fmt.Sprintf("File doesn't exists: %q \n", filepath)
}

func FileNotOpened(filepath string) string {
	return fmt.Sprintf("Cann't open file: %q \n", filepath)
}

func FileNotLoaded(filepath string) string {
	return fmt.Sprintf("Cann't load or parse file: %q \n", filepath)
}
