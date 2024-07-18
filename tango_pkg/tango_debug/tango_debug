package tango_debug

import (
	"fmt"
	"reflect"
)

func separator() {
	fmt.Println("--------------------------------")
}

func debugBannerTop() {
	fmt.Println("")
	separator()
	fmt.Println("> DEBUG: ")
}

func debugBannerBottom() {
	separator()
	fmt.Println("")
}

func String(name, value string) {

	debugBannerTop()
	fmt.Printf("%s: %s \n", name, value)
	debugBannerBottom()

}

func Struct(name string, input interface{}) {

	value := reflect.ValueOf(input)
	numFields := value.NumField()

	debugBannerTop()
	fmt.Println("Name: ", name)
	fmt.Printf("Number of fields: %d\n", numFields)

	structType := value.Type()

	for i := 0; i < numFields; i++ {
		field := structType.Field(i)
		fieldValue := value.Field(i)

		fmt.Printf("Field %d: %s (%s) = %v\n", i+1, field.Name, field.Type, fieldValue)
	}
	debugBannerBottom()
}
