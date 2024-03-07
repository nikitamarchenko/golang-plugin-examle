package main

import (
	"fmt"
	"time"

	"github.com/nikitamarchenko/golang-plugin-examle/types"
)

func SetUp(f types.PrinterFunc) {
	printFunc = f
}

var printFunc types.PrinterFunc

func Loop() {
	go func() {
		for i := 0; ; i++ {
			printFunc(fmt.Sprintf("%d loop", i))
			time.Sleep(time.Second)
		}
	}()
}
