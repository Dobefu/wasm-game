package game

import (
	"fmt"
	"math"
	"time"

	"github.com/Dobefu/wasm-game/cmd/canvas"
)

var (
	CANVAS     *canvas.Canvas
	FRAME      int
	DELTA_TIME int64
	_LAST_TIME time.Time
)

func init() {
	CANVAS = &canvas.CANVAS
	_LAST_TIME = time.Now()
}

func Update() {
	FRAME = FRAME + 1
}

func Render(clearCanvas bool) {
	DELTA_TIME = time.Now().UnixMilli() - _LAST_TIME.UnixMilli()
	_LAST_TIME = time.Now()
	fmt.Println(DELTA_TIME)

	if clearCanvas {
		CANVAS.Context.Call("clearRect", 0, 0, CANVAS.Width, CANVAS.Height)
	}

	sin := math.Sin(float64(FRAME) / 30)
	cos := math.Cos(float64(FRAME) / 30)

	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("moveTo", (float64(CANVAS.Width/2)*cos)+(float64(CANVAS.Width)/2), (float64(CANVAS.Height/2)*sin)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("lineTo", -(float64(CANVAS.Width/2)*cos)+(float64(CANVAS.Width)/2), -(float64(CANVAS.Height/2)*sin)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("stroke")
}
