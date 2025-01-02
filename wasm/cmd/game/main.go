package game

import (
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

	obj := Instantiate(structs.GameObject{Width: 100, Height: 100, X: 500, Y: 500})
	Instantiate(structs.GameObject{Width: 100, Height: 100, X: 150, Parent: obj})
	Instantiate(structs.GameObject{Width: 100, Height: 100, Y: 150, Parent: obj})
}

func Update() {
	FRAME = FRAME + 1
}

func Render(clearCanvas bool) {
	DELTA_TIME = float64(time.Now().UnixMilli()-_LAST_TIME.UnixMilli()) / 1000
	_LAST_TIME = time.Now()

	if clearCanvas {
		CANVAS.Context.Set("fillStyle", "black")
		CANVAS.Context.Call("fillRect", 0, 0, CANVAS.Width, CANVAS.Height)
	}

	for _, gameObject := range GAME_OBJECTS {
		rotation := gameObject.Rotation

		if gameObject.Parent == nil {
			gameObject.Rotation += DELTA_TIME * 50
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

		CANVAS.Context.Call("save")
		CANVAS.Context.Set("fillStyle", "white")
		CANVAS.Context.Call("translate", x, y)
		CANVAS.Context.Call("rotate", rotate.ToRadians(rotation))
		CANVAS.Context.Call("beginPath")
		CANVAS.Context.Call("rect", -(gameObject.Width / 2), -(gameObject.Height / 2), gameObject.Width, gameObject.Height)
		CANVAS.Context.Call("fill")
		CANVAS.Context.Call("restore")
	}
}
