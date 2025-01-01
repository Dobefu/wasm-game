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

	obj := Instantiate(structs.GameObject{X: 100, Y: 100})
	Instantiate(structs.GameObject{X: 250, Parent: obj})
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
		rotation := gameObject.Rotation

		var xFrom, yFrom, xTo, yTo float64

		if gameObject.Parent != nil {
			xFrom += gameObject.Parent.X
			yFrom += gameObject.Parent.Y
			xTo += gameObject.Parent.X
			yTo += gameObject.Parent.Y
			rotation += gameObject.Parent.Rotation
		}

		sin, cos := math.Sincos(rotation / (360 / (math.Pi * 2)))

		xFrom, yFrom = xFrom+gameObject.X-(cos*100), yFrom+gameObject.Y-(sin*100)
		xTo, yTo = xTo+gameObject.X+(cos*100), yTo+gameObject.Y+(sin*100)

		CANVAS.Context.Call("beginPath")
		CANVAS.Context.Call("moveTo", xFrom, yFrom)
		CANVAS.Context.Call("lineTo", xTo, yTo)
		CANVAS.Context.Call("stroke")
	}
}
