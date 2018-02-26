package builtin

import (
	"reflect"
)

func Int64(a interface{}) int64 {
	val := reflect.ValueOf(a)
	kind := val.Type().Kind()
	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return int64(val.Int())
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return int64(val.Uint())
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return int64(val.Float())
	}
	return int64(0)
}

func Uint64(a interface{}) uint64 {
	val := reflect.ValueOf(a)
	kind := val.Type().Kind()
	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return uint64(val.Int())
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return uint64(val.Uint())
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return uint64(val.Float())
	}
	return uint64(0)
}

func Int(a interface{}) int       { return int(Int64(a)) }
func Int8(a interface{}) int8     { return int8(Int64(a)) }
func Int16(a interface{}) int16   { return int16(Int64(a)) }
func Int32(a interface{}) int32   { return int32(Int64(a)) }
func Uint(a interface{}) uint     { return uint(Uint64(a)) }
func Uint8(a interface{}) uint8   { return uint8(Uint64(a)) }
func Uint16(a interface{}) uint16 { return uint16(Uint64(a)) }
func Uint32(a interface{}) uint32 { return uint32(Uint64(a)) }
