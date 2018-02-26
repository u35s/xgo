package xgo

import (
	"log"
)

var fntable = map[string]interface{}{
	"$ARITY":      (*Stack).Arity,
	"$push":       (*Stack).Push,
	"$ident":      (*Stack).PushIdent,
	"$arrayslice": (*Stack).PushArrayOrSlice,
	"$assign":     (*Stack).Assign,

	"$call":   (*XGo).Call,
	"$printf": log.Printf,
}

func Import(table map[string]interface{}) {
	for name, fn := range table {
		fntable[name] = fn
	}
}
