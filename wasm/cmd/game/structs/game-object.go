package structs

import "github.com/Dobefu/wasm-game/cmd/rotate"

type GameObject struct {
	X, Y     float64
	Width    float64
	Height   float64
	Scale    float64
	Rotation float64
	XSpeed   float64
	YSpeed   float64
	Parent   *GameObject
	Type     GameObjectType
}

func (obj *GameObject) GetWorldTransform() (x, y, rotation, scale float64) {
	x, y = obj.X, obj.Y
	rotation = obj.Rotation
	scale = obj.Scale

	if obj.Parent != nil {
		px, py, pr, ps := obj.Parent.GetWorldTransform()

		scaledOffsetX := obj.X * ps
		scaledOffsetY := obj.Y * ps

		rotatedX, rotatedY := rotate.RotatePoint(scaledOffsetX, scaledOffsetY, 0, 0, rotate.ToRadians(pr))

		x = px + rotatedX
		y = py + rotatedY
		rotation += pr
		scale *= ps
	}

	return x, y, rotation, scale
}
