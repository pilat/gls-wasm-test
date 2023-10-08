package main

//go:generate cp $GOROOT/misc/wasm/wasm_exec.js .

import (
	"fmt"
	"syscall/js"
	"time"

	"github.com/cosmos72/gls"
)

func main() {
	js.Global().Set("start", js.FuncOf(start))
	select {} // block forever
}

func start(this js.Value, args []js.Value) interface{} {
	print("button clicked")

	go f1()
	time.Sleep(500 * time.Millisecond)
	go f2()

	return nil
}

func f1() {
	for {
		print(fmt.Sprintf("f1(%d)", gls.GoID()))
		time.Sleep(1000 * time.Millisecond)
	}
}

func f2() {
	for {
		print(fmt.Sprintf("f2(%d)", gls.GoID()))
		time.Sleep(1000 * time.Millisecond)
	}
}

func print(s string) {
	fmt.Println(s)

	document := js.Global().Get("document")
	response := document.Call("getElementById", "response")
	response.Set("value", response.Get("value").String()+s+", ")
}
