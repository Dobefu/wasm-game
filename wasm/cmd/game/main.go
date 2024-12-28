package game

import (
	"math"

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

func Render(clearCanvas bool) {
	if clearCanvas {
		CANVAS.Context.Call("clearRect", 0, 0, CANVAS.Width, CANVAS.Height)
	}

	sin := math.Sin(float64(frame) / 30)
	cos := math.Cos(float64(frame) / 30)

	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("moveTo", (float64(CANVAS.Width/2)*cos)+(float64(CANVAS.Width)/2), (float64(CANVAS.Height/2)*sin)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("lineTo", -(float64(CANVAS.Width/2)*cos)+(float64(CANVAS.Width)/2), -(float64(CANVAS.Height/2)*sin)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("stroke")
}
