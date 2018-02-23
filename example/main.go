package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/u35s/xgo"

	"text/tpl.v1/interpreter"
)

type X struct {
	x1 struct {
		ra float64
		rb []int64
	}
	x2 float64
	X3 func() int64
	x4 func() int64
}

var x X = X{}

/////////////////////////

var (
	ipt    = xgo.New()
	engine *interpreter.Engine
)

func eval(line string) {

	defer func() {
		if err := recover(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}()

	line = strings.Trim(line, " \t\r\n")
	if line == "" {
		return
	}

	if err := engine.Eval(line); err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	v, _ := ipt.Ret()
	fmt.Printf("> %v\n\n", v)
}

func main() {
	x.x1.ra = 6
	x.x1.rb = []int64{1}
	x.X3 = func() int64 { return 5 }
	x.x4 = func() int64 { return 5 }
	ipt.AddVar("x", x)
	var err error
	if engine, err = interpreter.New(ipt, nil); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	code := `
1 + 3 * 5 + x.x1.rb[0] + add(1,7) + x.X3()
	`
	eval(code)
}
