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
	GAME_OBJECTS []*structs.GameObject
)

func init() {
	CANVAS = &canvas.CANVAS
	_LAST_TIME = time.Now()

	Instantiate(structs.GameObject{X: 300, Y: 100})
	Instantiate(structs.GameObject{X: 100, Y: 300})
	Instantiate(structs.GameObject{X: 100, Y: 100, Rotation: 90})
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

	for _, gameObject := range GAME_OBJECTS {
		sin, cos := math.Sincos(gameObject.Rotation / (360 / (math.Pi * 2)))

		CANVAS.Context.Call("beginPath")
		CANVAS.Context.Call("moveTo", gameObject.X+(cos*100), gameObject.Y+(sin*100))
		CANVAS.Context.Call("lineTo", gameObject.X-(cos*100), gameObject.Y-(sin*100))
		CANVAS.Context.Call("stroke")
	}
}
