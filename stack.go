package xgo

import (
	"fmt"
	"log"
	"reflect"
	"runtime"
	"strings"

	tpl "text/tpl.v1"
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

func (stk *Stack) getVal(n string) reflect.Value {
	slc := strings.Split(n, ".")
	v, ok := stk.vartable[slc[0]]
	if !ok {
		panic("undefined var: " + slc[0])
	}
	if len(slc) == 1 {
		if val := reflect.ValueOf(v); val.Kind() == reflect.Interface || val.Kind() == reflect.Ptr {
			return val.Elem()
		} else {
			return val
		}
	}
	_, val, ok := recursiveGetReflectValue(v, slc[1:])
	if !ok {
		panic("undefined var: " + n)
	}
	return val
}

func (stk *Stack) Assign(v interface{}) {
	temp, ok := stk.Pop()
	if !ok {
		panic("assign with stack empty")
	}
	slc := v.([]tpl.Token)
	v1 := slc[0].Literal
	if strings.Contains(v1, ".") {
		val := stk.getVal(v1)
		val.Set(reflect.ValueOf(temp))
	} else {
		stk.vartable[v1] = temp
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
		switch kind := v.Kind(); {
		case kind == reflect.Float64 || kind == reflect.Float32:
			stk.Push(v.Float())
		case kind >= reflect.Int && kind <= reflect.Int64:
			stk.Push(int(v.Int()))
		case kind >= reflect.Uint && kind <= reflect.Uintptr:
			stk.Push(uint(v.Uint()))
		case kind == reflect.Ptr:
			stk.Push(v.Interface())
		case kind == reflect.Interface:
			stk.Push(v.Interface())
		default:
			log.Println(v.Kind())
			return ErrUnsupportedRetType
		}
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

func (stk *Stack) Arity(arity int) {
	stk.Push(arity)
}
