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
}

func init() {
	xgo.Import(exports)
}
