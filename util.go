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
	for i := 0; i < val.NumMethod(); i++ {
		if val.Type().Method(i).Name == name {
			return i, val.Method(i), true
		}
	}
	return 0, val, false
}

func recursiveGetReflectValue(itfc interface{}, nameSlc []string) (index int, val reflect.Value, ok bool) {
	for j := 0; j < len(nameSlc); j++ {
		if j == 0 {
			val = reflect.ValueOf(itfc)
			if val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
				for i := 0; i < val.NumMethod(); i++ {
					if val.Type().Method(i).Name == nameSlc[j] {
						index, val, ok = i, val.Method(i), true
						return
					}
				}
				val = val.Elem()
			}
		}
		index, val, ok = getReflectValue(val, nameSlc[j])
		if !ok {
			return
		}
	}
	return
}

func Atoi(s string) int64 {
	if i, err := strconv.ParseInt(s, 10, 0); err == nil {
		return int64(i)
	}
	return 0
}

func Atof(s string) float64 {
	if i, err := strconv.ParseFloat(s, 10); err == nil {
		return float64(i)
	}
	return 0
}
