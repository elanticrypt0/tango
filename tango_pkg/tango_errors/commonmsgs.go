package tango_errors

import "fmt"

func MsgIDNotFound(id int) string {
	return fmt.Sprintf("Record not found with the id %d", id)
}

func MsgUIDNotFound(id string) string {
	return fmt.Sprintf("Record not found with the id %s", id)
}

func MsgNotFound() string {
	return "Record not found"
}

func MsgZeroRecordsFound() string {
	return "Zero records found"
}
