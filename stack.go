package xgo

import (
	"fmt"
	"log"
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
	if itfc, ok := v.([]tpl.Token); ok {
		switch itfc[0].Kind {
		case tpl.FLOAT:
			v1, _ := interpreter.ParseFloat(itfc[0].Literal)
			stk.stk = append(stk.stk, v1)
		case tpl.INT:
			v1, _ := interpreter.ParseInt(itfc[0].Literal)
			stk.stk = append(stk.stk, v1)
		}
	} else {
		stk.stk = append(stk.stk, v)
	}
}

func (stk *Stack) PushArrayOrSlice(v interface{}) {
	slc := v.([]tpl.Token)
	v1 := stk.getVal(slc[0].Literal)
	v2 := v1.Index(Atoi(slc[2].Literal))
	switch v2.Kind() {
	case reflect.Float64:
		stk.stk = append(stk.stk, v2.Float())
	}
}

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
	asignv, ok := stk.Pop()
	if !ok {
		panic("assign with stack empty")
	}
	slc := v.([]tpl.Token)
	v1 := slc[0].Literal
	if strings.Contains(v1, ".") {
		val := stk.getVal(v1)
		log.Printf("%+v,%+v,%+v,%+v,%v\n", v1, val, reflect.TypeOf(asignv).Kind(), val.Kind(), val.CanSet())
		switch kind := val.Kind(); {
		case kind == reflect.Bool:
			val.SetBool(asignv.(bool))
		case kind >= reflect.Int && kind <= reflect.Int64:
			val.SetInt(asignv.(int64))
		case kind >= reflect.Float32 && kind <= reflect.Float64:
			val.SetFloat(asignv.(float64))
		case kind >= reflect.Uint && kind <= reflect.Uintptr:
			if tp := reflect.TypeOf(asignv); tp.Kind() >= reflect.Float32 && tp.Kind() <= reflect.Float64 {
				val.SetUint(uint64(asignv.(float64)))
			} else if tp := reflect.TypeOf(asignv); tp.Kind() >= reflect.Uint && tp.Kind() <= reflect.Uint64 {
				val.SetUint(asignv.(uint64))
			}
		default:
			panic("unkonw assign type")
		}
	} else {
		stk.vartable[v1] = asignv
	}
}

func (stk *Stack) PushIdent(n string) {
	val := stk.getVal(n)
	switch kind := val.Kind(); {
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		stk.stk = append(stk.stk, val.Float())
	case kind >= reflect.Int && kind <= reflect.Int64:
		stk.stk = append(stk.stk, val.Int())
	case kind >= reflect.Uint && kind <= reflect.Uintptr:
		stk.stk = append(stk.stk, val.Uint())
	default:
		panic("error value:" + n + ":" + val.Kind().String())
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
		case kind == reflect.Ptr:
			stk.Push(v.Interface())
			continue
		default:
			log.Println(v.Kind())
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
