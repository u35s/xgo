package builtin

import "github.com/u35s/xgo"

var exports = map[string]interface{}{
	"$lshr": Lshr,
	"$rshr": Rshr,

	"$neg": Neg,
	"$add": Add,
	"$sub": Sub,
	"$mul": Mul,
	"$quo": Quo,
	"$mod": Mod,

	"$int":    Int,
	"$int8":   Int8,
	"$int16":  Int16,
	"$int32":  Int32,
	"$int64":  Int64,
	"$uint":   Uint,
	"$uint8":  Uint8,
	"$uint16": Uint16,
	"$uint32": Uint32,
}

func init() {
	xgo.Import(exports)
}
