package xgo

import (
	"errors"
	"fmt"
	"reflect"

	interpreter "text/tpl.v1/interpreter.util"
)

var (
	ErrUnsupportedRetType = errors.New("unsupported return type of function")
	ErrFncallWithoutArity = errors.New("function call without arity")
)

const gm = `

term = factor *('*' factor/mul | '/' factor/quo | '%' factor/mod)

doc = term *('+' term/add | '-' term/sub)

factor =
	FLOAT/push |
	STRING/push |
	'-' factor/neg |
	'(' doc ')' |
	(IDENT '(' doc %= ','/ARITY ')')/call |
	(IDENT '[' INT ']')/arrayslice |
	(IDENT '=' doc)/assign |
	IDENT/ident |
	'+' factor
`

type XGo struct {
	gm       string
	stack    *Stack
	fntable  map[string]interface{}
	vartable map[string]interface{}
}

func (x *XGo) Grammar() string                 { return x.gm }
func (x *XGo) Fntable() map[string]interface{} { return x.fntable }
func (x *XGo) Stack() interpreter.Stack        { return x.stack }

func New() *XGo {
	return &XGo{
		gm:       gm,
		stack:    NewStack(),
		fntable:  fntable,
		vartable: make(map[string]interface{}),
	}
}

func (x *XGo) Ret() (v interface{}, ok bool) {
	v, ok = x.stack.Pop()
	x.stack.Clear()
	return
}

func (x *XGo) Call(name string) error {
	var n int
	if arity, ok := x.stack.Pop(); ok {
		atp := reflect.TypeOf(arity)
		switch atp.Kind() {
		case reflect.Int:
			n = arity.(int)
		default:
			panic("unknow arity type " + atp.String())
		}
	} else {
		return ErrFncallWithoutArity
	}
	if fn, ok := x.fntable["$"+name]; ok {
		err := interpreter.Call(x.stack, fn, n)
		if err != nil {
			return fmt.Errorf("call function `%s` failed: %v", name, err)
		}
		return nil
	} else {
		val := x.stack.getVal(name)
		err := interpreter.Call(x.stack, val.Interface(), n)
		if err != nil {
			return fmt.Errorf("call function `%s` failed: %v", name, err)
		}
		return nil
	}
	return fmt.Errorf("function `%s` not found", name)
}

func (x *XGo) AddVar(n string, v interface{}) { x.stack.vartable[n] = v }
