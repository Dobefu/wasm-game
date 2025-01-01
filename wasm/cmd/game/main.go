package game

import (
	"math"
	"time"

	"github.com/Dobefu/wasm-game/cmd/canvas"
	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

var (
	CANVAS       *canvas.Canvas
	FRAME        int
	DELTA_TIME   float64
	_LAST_TIME   time.Time
	ROTATION     float64
	GAME_OBJECTS []structs.GameObject
)

func init() {
	CANVAS = &canvas.CANVAS
	_LAST_TIME = time.Now()
}

func Update() {
	FRAME = FRAME + 1
}

func Render(clearCanvas bool) {
	DELTA_TIME = float64(time.Now().UnixMilli()-_LAST_TIME.UnixMilli()) / 1000
	_LAST_TIME = time.Now()

	if clearCanvas {
		CANVAS.Context.Call("clearRect", 0, 0, CANVAS.Width, CANVAS.Height)
	}

	ROTATION += DELTA_TIME
	sin := math.Sin(ROTATION)
	cos := math.Cos(ROTATION)

	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("moveTo", (cos*100)+(float64(CANVAS.Width)/2), (sin*100)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("lineTo", -(cos*100)+(float64(CANVAS.Width)/2), -(sin*100)+(float64(CANVAS.Height)/2))
	CANVAS.Context.Call("stroke")
}
