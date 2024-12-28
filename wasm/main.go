//go:build js && wasm
// +build js,wasm

package main

import (
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/dom"
)

var (
	CANVAS  js.Value
	CONTEXT js.Value
)

func init() {
	CANVAS = dom.GetElementById("game")
	CONTEXT = CANVAS.Call("getContext", "2d")

	dom.AddEventListener("window", "resize", func() { resizeCanvas(true) })
}

func main() {
	resizeCanvas(false)
	render()

	ch := make(chan struct{})
	<-ch
}

func render() {
	for i := 0; i < 50; i++ {
		CONTEXT.Call("beginPath")
		CONTEXT.Call("moveTo", 0, 0)
		CONTEXT.Call("lineTo", 100, 100)
		CONTEXT.Call("stroke")
	}
}

func resizeCanvas(rerender bool) {
	CANVAS.Set("height", dom.WINDOW.Get("innerHeight"))
	CANVAS.Set("width", dom.WINDOW.Get("innerWidth"))

	if rerender {
		render()
	}
}
