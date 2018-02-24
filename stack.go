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

func (stk *Stack) PushArrayOrSlice(v interface{}) {
	slc := v.([]tpl.Token)
	v1 := stk.getVal(slc[0].Literal)
	v2 := v1.Index(int(Atoi(slc[2].Literal)))
	switch v2.Kind() {
	case reflect.Float64:
		stk.stk = append(stk.stk, v2.Float())
	case reflect.Int:
		stk.stk = append(stk.stk, int(v2.Int()))
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
		log.Printf("val:%+v,%+v,%+v,%+v,assign:%v\n", v1, val, val.Kind(), val.CanSet(), reflect.TypeOf(asignv).Kind())
		switch kind := val.Kind(); {
		case kind == reflect.Bool:
			val.SetBool(asignv.(bool))
		case kind >= reflect.Int && kind <= reflect.Int64:
			val.SetInt(asignv.(int64))
		case kind >= reflect.Float32 && kind <= reflect.Float64:
			val.SetFloat(asignv.(float64))
		case kind >= reflect.Uint && kind <= reflect.Uintptr:
			atp := reflect.TypeOf(asignv)
			if atp.Kind() >= reflect.Float32 && atp.Kind() <= reflect.Float64 {
				val.SetUint(uint64(asignv.(float64)))
			} else if atp.Kind() >= reflect.Uint && atp.Kind() <= reflect.Uint64 {
				val.SetUint(asignv.(uint64))
			} else if atp.Kind() >= reflect.Int && atp.Kind() <= reflect.Int64 {
				val.SetUint(uint64(asignv.(int64)))
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
	case kind == reflect.Struct:
		stk.stk = append(stk.stk, val.Interface())
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
		case kind == reflect.Interface:
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
