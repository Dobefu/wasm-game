package game

import (
	"github.com/Dobefu/wasm-game/cmd/canvas"
)

var (
	CANVAS *canvas.Canvas
	frame  int64
)

func init() {
	CANVAS = &canvas.CANVAS
}

func Update() {
	frame = frame + 1
}

func Render() {
	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("moveTo", 0, 0)
	CANVAS.Context.Call("lineTo", CANVAS.Width, CANVAS.Height)
	CANVAS.Context.Call("stroke")

	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("moveTo", CANVAS.Width, 0)
	CANVAS.Context.Call("lineTo", 0, CANVAS.Height)
	CANVAS.Context.Call("stroke")
}
