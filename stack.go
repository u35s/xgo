package xgo

import (
	"fmt"
	"reflect"
	"runtime"
	"strings"

	tpl "text/tpl.v1"
	"text/tpl.v1/interpreter"
)

func printStack(all bool) {
	bs := make([]byte, 1<<18)
	cnt := string(bs[:runtime.Stack(bs, all)])
	fmt.Println(cnt)
}

type Stack struct {
	stk      []interface{}
	vartable map[string]interface{}
}

func NewStack() *Stack    { return &Stack{vartable: make(map[string]interface{})} }
func (stk *Stack) Clear() { stk.stk = (stk.stk)[:0] }
func (stk *Stack) Push(v interface{}) {
	tp := reflect.TypeOf(v)
	switch tp.Kind() {
	case reflect.Slice:
		itfc, ok := v.([]tpl.Token)
		if !ok {
			panic("unknow slice type")
		}
		switch itfc[0].Kind {
		case tpl.FLOAT:
			v1, _ := interpreter.ParseFloat(itfc[0].Literal)
			stk.stk = append(stk.stk, v1)
		case tpl.INT:
			v1, _ := interpreter.ParseInt(itfc[0].Literal)
			stk.stk = append(stk.stk, v1)
		}
	case reflect.Float32:
		stk.stk = append(stk.stk, v)
	case reflect.Float64:
		stk.stk = append(stk.stk, v)
	case reflect.Int:
		stk.stk = append(stk.stk, v)
	case reflect.Int32:
		stk.stk = append(stk.stk, v)
	case reflect.Int64:
		stk.stk = append(stk.stk, v)
	default:
		panic("unknow push type : " + tp.String())
	}
}

func (stk *Stack) PushArrayOrSlice(v interface{}) {
	slc := v.([]tpl.Token)
	v1 := stk.getVal(slc[0].Literal)
	v2 := v1.Index(Atoi(slc[2].Literal))
	switch v2.Kind() {
	case reflect.Int64:
		stk.stk = append(stk.stk, int(v2.Int()))
	}
}

func (stk *Stack) getVal(n string) reflect.Value {
	slc := strings.Split(n, ".")
	v, ok := stk.vartable[slc[0]]
	if !ok {
		panic("undefined var: " + slc[0])
	}
	_, val, ok := recursiveGetReflectValue(v, slc[1:])
	if !ok {
		panic("undefined var: " + n)
	}
	return val
}

func (stk *Stack) PushIdent(n string) {
	val := stk.getVal(n)
	switch val.Kind() {
	case reflect.Float32:
		stk.stk = append(stk.stk, val.Float())
	case reflect.Float64:
		stk.stk = append(stk.stk, val.Float())
	case reflect.Interface:
		stk.stk = append(stk.stk, val.Interface())
	default:
		panic("error value:" + n)
	}
}

func (stk *Stack) Pop() (v interface{}, ok bool) {
	n := len(stk.stk)
	if n > 0 {
		v, ok = (stk.stk)[n-1], true
		stk.stk = (stk.stk)[:n-1]
	}
	return
}

func (stk *Stack) PushRet(ret []reflect.Value) error {
	for _, v := range ret {
		var val float64
		switch kind := v.Kind(); {
		case kind == reflect.Float64 || kind == reflect.Float32:
			val = v.Float()
		case kind >= reflect.Int && kind <= reflect.Int64:
			val = float64(v.Int())
		case kind >= reflect.Uint && kind <= reflect.Uintptr:
			val = float64(v.Uint())
		default:
			return ErrUnsupportedRetType
		}
		stk.Push(val)
	}
	return nil
}

func (stk *Stack) PopArgs(arity int) (args []reflect.Value, ok bool) {
	pstk := stk.stk
	n := len(pstk)
	if n >= arity {
		args, ok = make([]reflect.Value, arity), true
		n -= arity
		for i := 0; i < arity; i++ {
			args[i] = reflect.ValueOf(pstk[n+i])
		}
		stk.stk = pstk[:n]
	}
	return
}

func Arity(stk *Stack, arity int) {
	stk.Push(arity)
}
