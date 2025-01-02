package game

import (
	"math"
	"time"

	"github.com/Dobefu/wasm-game/cmd/canvas"
	"github.com/Dobefu/wasm-game/cmd/game/structs"
	"github.com/Dobefu/wasm-game/cmd/rotate"
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

	obj := Instantiate(structs.GameObject{X: 500, Y: 500})
	Instantiate(structs.GameObject{X: 150, Parent: obj})
	Instantiate(structs.GameObject{Y: 150, Parent: obj})
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

		if gameObject.Parent == nil {
			gameObject.Rotation += 1
		}

		var x, y float64

		if gameObject.Parent != nil {
			parentRotation := rotate.ToRadians(gameObject.Parent.Rotation)
			rotatedX, rotatedY := rotate.RotatePoint(gameObject.X, gameObject.Y, 0, 0, parentRotation)

			x = gameObject.Parent.X + rotatedX
			y = gameObject.Parent.Y + rotatedY

			rotation = gameObject.Parent.Rotation + rotation
		} else {
			x = gameObject.X
			y = gameObject.Y
		}

		radians := rotate.ToRadians(rotation)
		sin, cos := math.Sincos(radians)

		xFrom := x - cos*50
		yFrom := y - sin*50
		xTo := x + cos*50
		yTo := y + sin*50

		CANVAS.Context.Call("beginPath")
		CANVAS.Context.Call("moveTo", xFrom, yFrom)
		CANVAS.Context.Call("lineTo", xTo, yTo)
		CANVAS.Context.Call("stroke")
	}
}
