//go:build js && wasm
// +build js,wasm

package main

import (
	"fmt"
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/dom"
)

func init() {
	dom.AddEventListener("window", "resize", func() {
		fmt.Println("Resize")
	})
}

func main() {
	render()

	ch := make(chan struct{})
	<-ch
}

func render() {
	var canvas js.Value = js.Global().Get("document").Call("getElementById", "game")
	var context js.Value = canvas.Call("getContext", "2d")

	canvas.Set("height", 100)
	canvas.Set("width", 100)

	for i := 0; i < 50; i++ {
		context.Call("beginPath")
		context.Call("moveTo", 0, 0)
		context.Call("lineTo", 100, 100)
		context.Call("stroke")
	}
}
