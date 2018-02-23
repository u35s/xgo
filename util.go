package xgo

import (
	"reflect"
	"strconv"
)

func getReflectValue(val reflect.Value, name string) (int, reflect.Value, bool) {
	for i := 0; i < val.NumField(); i++ {
		if val.Type().Field(i).Name == name {
			return i, val.Field(i), true
		}
	}
	return 0, val, false
}

func recursiveGetReflectValue(itfc interface{}, nameSlc []string) (index int, val reflect.Value, ok bool) {
	for j := 0; j < len(nameSlc); j++ {
		if j == 0 {
			val = reflect.ValueOf(itfc)
		}
		index, val, ok = getReflectValue(val, nameSlc[j])
		if !ok {
			return
		}
	}
	return
}

func Atoi(s string) int {
	if i, err := strconv.ParseInt(s, 10, 0); err == nil {
		return int(i)
	}
	return 0
}
