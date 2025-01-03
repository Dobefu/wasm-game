//go:build js && wasm
// +build js,wasm

package main

import (
	"github.com/Dobefu/wasm-game/cmd/dom"
	"github.com/Dobefu/wasm-game/cmd/game"
)

func main() {
	dom.RequestAnimationFrame(func() {
		game.Update()
		game.Render(true)
	})

	ch := make(chan struct{})
	<-ch
}
