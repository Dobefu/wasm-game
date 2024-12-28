//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/Dobefu/wasm-game/cmd/canvas"
	"github.com/Dobefu/wasm-game/cmd/dom"
)

func main() {
	dom.RequestAnimationFrame(func() {
		canvas.Render()
	})

	ch := make(chan struct{})
	<-ch
}
