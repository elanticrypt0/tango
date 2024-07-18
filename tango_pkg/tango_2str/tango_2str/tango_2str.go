package tango_2str

import (
	"fmt"
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
