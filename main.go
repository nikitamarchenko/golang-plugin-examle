package main

import (
	"fmt"
	"plugin"
	"time"

	"github.com/nikitamarchenko/golang-plugin-examle/types"
)

func PrintFromMain(s string) {
	fmt.Printf("main app: %s\n", s)
}

func main() {
	p, err := plugin.Open("./plugins/hello/hello.so")

	setUp, err := p.Lookup("SetUp")
	if err != nil {
		panic(err)
	}
	setUp.(func(types.PrinterFunc))(PrintFromMain)

	loop, err := p.Lookup("Loop")
	if err != nil {
		panic(err)
	}

	loop.(func())()
	time.Sleep(time.Second * 20)
}
