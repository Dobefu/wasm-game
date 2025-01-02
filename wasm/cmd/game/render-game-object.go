package game

import (
	"github.com/Dobefu/wasm-game/cmd/game/structs"
	"github.com/Dobefu/wasm-game/cmd/rotate"
)

func RenderGameObject(gameObject *structs.GameObject) {
	rotation := gameObject.Rotation
	scale := gameObject.Scale

	if gameObject.Parent == nil {
		gameObject.Rotation += DELTA_TIME * 50
	}

	var x, y float64

	if gameObject.Parent != nil {
		parentRotation := rotate.ToRadians(gameObject.Parent.Rotation)
		rotatedX, rotatedY := rotate.RotatePoint(gameObject.X*gameObject.Parent.Scale, gameObject.Y*gameObject.Parent.Scale, 0, 0, parentRotation)

		x = gameObject.Parent.X + rotatedX
		y = gameObject.Parent.Y + rotatedY

		rotation += gameObject.Parent.Rotation
		scale *= gameObject.Parent.Scale
	} else {
		x = gameObject.X
		y = gameObject.Y
	}

	width := gameObject.Width * scale
	height := gameObject.Height * scale

	CANVAS.Context.Call("save")
	CANVAS.Context.Set("fillStyle", "white")
	CANVAS.Context.Call("translate", x, y)
	CANVAS.Context.Call("rotate", rotate.ToRadians(rotation))
	CANVAS.Context.Call("beginPath")
	CANVAS.Context.Call("rect", -width/2, -height/2, width, height)
	CANVAS.Context.Call("fill")
	CANVAS.Context.Call("restore")
}
