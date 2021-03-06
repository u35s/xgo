package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/u35s/xgo"

	_ "github.com/u35s/xgo/lib/builtin"

	"text/tpl.v1/interpreter"
)

type X struct {
	x1 struct {
		ra int
		rb []int
	}
	x2 int
	X3 func(int) int
	x4 func() int
	X5 int
}

type Y struct {
	Y1 int
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
	x.x1.rb = []int{1}
	x.X3 = func(d int) int { return 5 * d }
	x.x4 = func() int { return 5 }
	x.X5 = 3
	ipt.AddVar("x", &x)
	var err error
	if engine, err = interpreter.New(ipt, nil); err != nil {
		log.Printf("%v\n", err)
		os.Exit(1)
	}
	code := `
temp = 1 + 3 * 5 + x.x1.rb[0] + add(1,7) + x.X3(x.X5)
printf("temp:%v",temp)
temp = temp + int(100)
printf("temp:%v",temp)
x.X5 = temp + 60
printf("x.X5:%v",x.X5)
y1 = x.Get()
y1.Y1 = x.X5
printf("x:%+v,y1:%+v\n",x,y1)

a = 1<<5
printf("a:%+v",a)
b = 1
a = a>>b
printf("a:%+v,b:%+v",a,b)

u1 = uint(1)+2
printf("u1:%+v",u1)
`
	for _, v := range strings.Split(code, "\n") {
		eval(v)
	}
}
