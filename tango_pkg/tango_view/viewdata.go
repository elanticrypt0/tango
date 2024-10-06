package tango_view

import (
	"fmt"
	"strconv"
	"tango_pkg/tango_log"
)

type ViewData struct {
	Data      map[string]interface{}
	DataSlice map[string][]interface{}
}

func NewViewData() ViewData {
	return ViewData{
		Data:      map[string]interface{}{},
		DataSlice: map[string][]interface{}{},
	}
}

func (me *ViewData) Set(name string, data interface{}) {
	me.Data[name] = data
}

func (me *ViewData) SetSlice(name string, data []interface{}) {
	me.DataSlice[name] = data
}

func (me *ViewData) Get(name string) interface{} {
	if me.IsInData(name) {
		return me.Data[name]
	} else {
		return nil
	}
}

func (me *ViewData) GetSlice(name string) []interface{} {
	if me.IsInDataSlice(name) {
		return me.DataSlice[name]
	} else {
		return nil
	}
}

func (me *ViewData) GetSliceAsStrings(name string) []string {
	sliceAux := []string{}
	if me.IsInDataSlice(name) {
		for _, val := range me.DataSlice[name] {
			aux := fmt.Sprintf("%v", val)
			sliceAux = append(sliceAux, aux)
		}
		return sliceAux
	} else {
		return nil
	}
}

func (me *ViewData) GetAsString(name string) string {
	if me.IsInData(name) {
		return fmt.Sprint(me.Data[name])
	}
	return ""
}

func (me *ViewData) GetAsInt(name string) int {
	if me.IsInData(name) {
		val := me.GetAsString(name)
		valAux, err := strconv.Atoi(val)
		if err != nil {
			tango_log.PrintError(err.Error())
		}
		return valAux
	}
	return 0
}

func (me *ViewData) GetAll() map[string]interface{} {
	return me.Data
}

func (me *ViewData) GetAllSlices() map[string][]interface{} {
	return me.DataSlice
}

func (me *ViewData) IsInData(name string) bool {
	finded := false
	for key := range me.Data {
		if name == key {
			finded = true
		}
	}
	return finded
}

func (me *ViewData) IsInDataSlice(name string) bool {
	finded := false
	for key := range me.DataSlice {
		if name == key {
			finded = true
		}
	}
	return finded
}
