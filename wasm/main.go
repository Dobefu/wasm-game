//go:build js && wasm
// +build js,wasm

package main

import "github.com/Dobefu/wasm-game/cmd/canvas"

func main() {
	canvas.Render()

	ch := make(chan struct{})
	<-ch
}
