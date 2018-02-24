package builtin

import "reflect"

func panicUnsupportedOp1(op string, a interface{}) interface{} {
	ta := typeString(a)
	panic("unsupported operator: " + op + ta)
}

func panicUnsupportedOp2(op string, a, b interface{}) interface{} {

	ta := typeString(a)
	tb := typeString(b)
	panic("unsupported operator: " + ta + op + tb)
}

func typeString(a interface{}) string {

	if a == nil {
		return "nil"
	}
	return reflect.TypeOf(a).String()
}
