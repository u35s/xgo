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
		rb []float64
	}
	x2 float64
	X3 func(uint) float64
	x4 func() float64
	X5 uint
}

type Y struct {
	Y1 uint
}

func (this *X) Get() *Y { return &y }

var x X = X{}

var y Y = Y{}

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
}

func main() {
	x.x1.ra = 6
	x.x1.rb = []float64{1}
	x.X3 = func(d uint) float64 { return 5 * float64(d) }
	x.x4 = func() float64 { return 5 }
	x.X5 = 3
	ipt.AddVar("x", &x)
	var err error
	if engine, err = interpreter.New(ipt, nil); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	code := `
temp = 1 + 3 * 5 + x.x1.rb[0] + add(1,7) + x.X3(x.X5)
temp = temp + 100
x.X5 = temp + 60
y1 = x.Get()
y1.Y1 = x.X5
logi("x:%+v,y1:%+v\n",x,y1)
`
	for _, v := range strings.Split(code, "\n") {
		eval(v)
	}
}
