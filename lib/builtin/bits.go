package builtin

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
