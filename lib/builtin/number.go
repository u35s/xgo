package builtin

// Neg returns -a
func Neg(a interface{}) interface{} {

	switch a1 := a.(type) {
	case int64:
		return -a1
	case float64:
		return -a1
	}
	return panicUnsupportedOp1("-", a)
}

// Add returns a+b
func Add(a, b interface{}) interface{} {
	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 + b1
		case float64:
			return float64(a1) + b1
		}
	case float64:
		switch b1 := b.(type) {
		case int:
			return a1 + float64(b1)
		case float64:
			return a1 + b1
		}
	case string:
		if b1, ok := b.(string); ok {
			return a1 + b1
		}
	case uint:
		switch b1 := b.(type) {
		case int:
			return a1 + uint(b1)
		}
	case uint64:
		switch b1 := b.(type) {
		case int:
			return a1 + uint64(b1)
		}
	case int64:
		switch b1 := b.(type) {
		case int64:
			return a1 + int64(b1)
		case float64:
			return a1 + int64(b1)
		}
	case uint32:
		switch b1 := b.(type) {
		case int:
			return a1 + uint32(b1)
		}
	case int32:
		switch b1 := b.(type) {
		case int:
			return a1 + int32(b1)
		}
	case uint16:
		switch b1 := b.(type) {
		case int:
			return a1 + uint16(b1)
		}
	case int16:
		switch b1 := b.(type) {
		case int:
			return a1 + int16(b1)
		}
	case uint8:
		switch b1 := b.(type) {
		case int:
			return a1 + uint8(b1)
		}
	case int8:
		switch b1 := b.(type) {
		case int:
			return a1 + int8(b1)
		}
	}
	return panicUnsupportedOp2("+", a, b)
}

// Sub returns a-b
func Sub(a, b interface{}) interface{} {
	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 - b1
		case float64:
			return float64(a1) - b1
		}
	case float64:
		switch b1 := b.(type) {
		case int:
			return a1 - float64(b1)
		case float64:
			return a1 - b1
		}
	case uint:
		switch b1 := b.(type) {
		case int:
			return a1 - uint(b1)
		}
	case uint64:
		switch b1 := b.(type) {
		case int:
			return a1 - uint64(b1)
		}
	case int64:
		switch b1 := b.(type) {
		case int:
			return a1 - int64(b1)
		}
	case uint32:
		switch b1 := b.(type) {
		case int:
			return a1 - uint32(b1)
		}
	case int32:
		switch b1 := b.(type) {
		case int:
			return a1 - int32(b1)
		}
	case uint16:
		switch b1 := b.(type) {
		case int:
			return a1 - uint16(b1)
		}
	case int16:
		switch b1 := b.(type) {
		case int:
			return a1 - int16(b1)
		}
	case uint8:
		switch b1 := b.(type) {
		case int:
			return a1 - uint8(b1)
		}
	case int8:
		switch b1 := b.(type) {
		case int:
			return a1 - int8(b1)
		}
	}
	return panicUnsupportedOp2("-", a, b)
}

// Mul returns a*b
func Mul(a, b interface{}) interface{} {
	switch a1 := a.(type) {
	case int64:
		switch b1 := b.(type) {
		case int64:
			return a1 * b1
		case float64:
			return float64(a1) * b1
		}
	case float64:
		switch b1 := b.(type) {
		case int64:
			return a1 * float64(b1)
		case float64:
			return a1 * b1
		}
	}
	return panicUnsupportedOp2("*", a, b)
}

// Quo returns a/b
func Quo(a, b interface{}) interface{} {
	switch a1 := a.(type) {
	case int:
		switch b1 := b.(type) {
		case int:
			return a1 / b1
		case float64:
			return float64(a1) / b1
		}
	case float64:
		switch b1 := b.(type) {
		case int:
			return a1 / float64(b1)
		case float64:
			return a1 / b1
		}
	}
	return panicUnsupportedOp2("/", a, b)
}

// Mod returns a%b
func Mod(a, b interface{}) interface{} {
	if a1, ok := a.(int); ok {
		if b1, ok := b.(int); ok {
			return a1 % b1
		}
	}
	return panicUnsupportedOp2("%", a, b)
}
