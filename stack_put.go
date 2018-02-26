package xgo

import (
	"reflect"

	tpl "text/tpl.v1"
)

func (stk *Stack) Push(v interface{}) {
	if itfc, ok := v.([]tpl.Token); ok {
		switch itfc[0].Kind {
		case tpl.FLOAT:
			stk.stk = append(stk.stk, Atof(itfc[0].Literal))
		case tpl.INT:
			stk.stk = append(stk.stk, Atoi(itfc[0].Literal))
		case tpl.STRING:
			stk.stk = append(stk.stk, itfc[0].Literal)
		}
	} else {
		stk.stk = append(stk.stk, v)
	}
}

func (stk *Stack) PushIdent(n string) {
	stk.PushValue(stk.getVal(n))
}

func (stk *Stack) PushArrayOrSlice(v interface{}) {
	slc := v.([]tpl.Token)
	stk.PushValue(stk.getVal(slc[0].Literal).Index(Atoi(slc[2].Literal)))
}

func (stk *Stack) PushValue(v reflect.Value) {
	switch v.Kind() {
	case reflect.Bool:
		stk.stk = append(stk.stk, v.Bool())
	case reflect.Int:
		stk.stk = append(stk.stk, int(v.Int()))
	case reflect.Int8:
		stk.stk = append(stk.stk, int8(v.Int()))
	case reflect.Int16:
		stk.stk = append(stk.stk, int16(v.Int()))
	case reflect.Int32:
		stk.stk = append(stk.stk, int32(v.Int()))
	case reflect.Int64:
		stk.stk = append(stk.stk, int64(v.Int()))
	case reflect.Uint:
		stk.stk = append(stk.stk, uint(v.Uint()))
	case reflect.Uint8:
		stk.stk = append(stk.stk, uint8(v.Uint()))
	case reflect.Uint16:
		stk.stk = append(stk.stk, uint16(v.Uint()))
	case reflect.Uint32:
		stk.stk = append(stk.stk, uint32(v.Uint()))
	case reflect.Uint64:
		stk.stk = append(stk.stk, uint64(v.Uint()))
	case reflect.Uintptr:
		stk.stk = append(stk.stk, uint64(v.Uint()))
	case reflect.Float32:
		stk.stk = append(stk.stk, float32(v.Float()))
	case reflect.Float64:
		stk.stk = append(stk.stk, float64(v.Float()))
	case reflect.Complex64:
		stk.stk = append(stk.stk, complex64(v.Complex()))
	case reflect.Complex128:
		stk.stk = append(stk.stk, complex128(v.Complex()))
	case reflect.Array:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Chan:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Func:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Interface:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Map:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Ptr:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.Slice:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.String:
		stk.stk = append(stk.stk, v.String())
	case reflect.Struct:
		stk.stk = append(stk.stk, v.Interface())
	case reflect.UnsafePointer:
		stk.stk = append(stk.stk, v.Interface())
	default:
		panic("push array or slice with unkonwed kind:" + v.Kind().String())
	}
}
