package tango_2str

import (
	"fmt"
	"strings"
)

func FromInt(num int) string {
	return fmt.Sprintf("%d", num)
}

func FromUint(num uint) string {
	return fmt.Sprintf("%d", num)
}

func FromBool(val bool) string {
	if val {
		return "1"
	} else {
		return "0"
	}
}

func FromBoolFull(val bool) string {
	return fmt.Sprintf("%v", val)
}

func ToBool(str string) bool {
	str = strings.ToLower(str)
	switch str {
	case "1":
		return true
	case "true":
		return true
	case "0":
		return false
	case "false":
		return false
	default:
		return false
	}
}
