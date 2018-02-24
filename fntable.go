package xgo

import (
	"log"
	"math"
)

func Neg(a float64) float64    { return -a }
func Mul(a, b float64) float64 { return a * b }
func Quo(a, b float64) float64 { return a / b }
func Add(a, b float64) float64 { return a + b }
func Sub(a, b float64) float64 { return a - b }
func Inf(a float64) float64 {
	v := 0
	if a < 0 {
		v = -1
	}
	return math.Inf(v)
}
func Jn(a, b float64) float64    { return math.Jn(int(a), b) }
func Yn(a, b float64) float64    { return math.Yn(int(a), b) }
func Ldexp(a, b float64) float64 { return math.Ldexp(a, int(b)) }
func Pow10(a float64) float64    { return math.Pow10(int(a)) }
func Max(args ...float64) (max float64) {
	if len(args) == 0 {
		return
	}
	max = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] > max {
			max = args[i]
		}
	}
	return
}
func Min(args ...float64) (min float64) {
	if len(args) == 0 {
		return
	}
	min = args[0]
	for i := 1; i < len(args); i++ {
		if args[i] < min {
			min = args[i]
		}
	}
	return
}

// Lshr returns a << b
func Lshr(a, b interface{}) interface{} {
	switch a1 := a.(type) {
	case int64:
		switch b1 := b.(type) {
		case int64:
			return a1 << uint(b1)
		}
	}
	return panicUnsupportedOp2("<<", a, b)
}

// Rshr returns a >> b
func Rshr(a, b interface{}) interface{} {

	switch a1 := a.(type) {
	case int64:
		switch b1 := b.(type) {
		case int64:
			return a1 >> uint(b1)
		}
	}
	return panicUnsupportedOp2(">>", a, b)
}

// Xor returns a ^ b
func Xor(a, b interface{}) interface{} {

	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 ^ b1
		}
	}
	return panicUnsupportedOp2("^", a, b)
}

// BitAnd returns a & b
func BitAnd(a, b interface{}) interface{} {

	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 & b1
		}
	}
	return panicUnsupportedOp2("&", a, b)
}

// BitOr returns a | b
func BitOr(a, b interface{}) interface{} {

	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 | b1
		}
	}
	return panicUnsupportedOp2("|", a, b)
}

// BitNot returns ^a
func BitNot(a interface{}) interface{} {

	switch a1 := a.(type) {
	case int:
		return ^a1
	}
	return panicUnsupportedOp1("^", a)
}

// AndNot returns a &^ b
func AndNot(a, b interface{}) interface{} {

	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 &^ b1
		}
	}
	return panicUnsupportedOp2("&^", a, b)
}

var fntable = map[string]interface{}{
	"$lshr":   Lshr,
	"$rshr":   Rshr,
	"$bitand": BitAnd,
	"$bitnot": BitNot,
	"$andnot": AndNot,

	"abs":       math.Abs,
	"acos":      math.Acos,
	"acosh":     math.Acosh,
	"asin":      math.Asin,
	"asinh":     math.Asinh,
	"atan":      math.Atan,
	"atan2":     math.Atan2,
	"atanh":     math.Atanh,
	"cbrt":      math.Cbrt,
	"ceil":      math.Ceil,
	"copysign":  math.Copysign,
	"cos":       math.Cos,
	"cosh":      math.Cosh,
	"dim":       math.Dim,
	"erf":       math.Erf,
	"erfc":      math.Erfc,
	"exp":       math.Exp,
	"exp2":      math.Exp2,
	"expm1":     math.Expm1,
	"floor":     math.Floor,
	"gamma":     math.Gamma,
	"hypot":     math.Hypot,
	"inf":       Inf,
	"j0":        math.J0,
	"j1":        math.J1,
	"jn":        Jn,
	"ldexp":     Ldexp,
	"ln":        math.Log,
	"log":       math.Log,
	"log10":     math.Log10,
	"log1p":     math.Log1p,
	"log2":      math.Log2,
	"logb":      math.Logb,
	"max":       Max,
	"min":       Min,
	"mod":       math.Mod,
	"NaN":       math.NaN,
	"nextafter": math.Nextafter,
	"pow":       math.Pow,
	"pow10":     Pow10,
	"remainder": math.Remainder,
	"sin":       math.Sin,
	"sinh":      math.Sinh,
	"sqrt":      math.Sqrt,
	"tan":       math.Tan,
	"tanh":      math.Tanh,
	"trunc":     math.Trunc,
	"y0":        math.Y0,
	"y1":        math.Y1,
	"yn":        Yn,

	"$neg":        Neg,
	"$mul":        Mul,
	"$quo":        Quo,
	"$mod":        math.Mod,
	"$add":        Add,
	"$sub":        Sub,
	"$ARITY":      Arity,
	"$push":       (*Stack).Push,
	"$ident":      (*Stack).PushIdent,
	"$arrayslice": (*Stack).PushArrayOrSlice,
	"$assign":     (*Stack).Assign,
	"$call":       (*XGo).Call,
	"$logi":       log.Printf,
}
