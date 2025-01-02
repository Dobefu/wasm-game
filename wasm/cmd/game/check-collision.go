package game

import (
	"math"

	"github.com/Dobefu/wasm-game/cmd/game/structs"
	"github.com/Dobefu/wasm-game/cmd/rotate"
)

func CheckCollision(source, target *structs.GameObject) bool {
	x1, y1, r1, s1 := source.GetWorldTransform()
	x2, y2, r2, s2 := target.GetWorldTransform()

	corners1 := rotate.GetRotatedCorners(x1, y1, source.Width*s1, source.Height*s1, r1)
	corners2 := rotate.GetRotatedCorners(x2, y2, target.Width*s2, target.Height*s2, r2)

	minX1, minY1, maxX1, maxY1 := getBoundingBox(corners1)
	minX2, minY2, maxX2, maxY2 := getBoundingBox(corners2)

	if !(maxX1 < minX2 || maxX2 < minX1 || maxY1 < minY2 || maxY2 < minY1) {
		return true
	}

	return false
}

func getBoundingBox(points [][2]float64) (minX, minY, maxX, maxY float64) {
	minX, minY = math.Inf(1), math.Inf(1)
	maxX, maxY = math.Inf(-1), math.Inf(-1)

	for _, p := range points {
		minX = math.Min(minX, p[0])
		minY = math.Min(minY, p[1])
		maxX = math.Max(maxX, p[0])
		maxY = math.Max(maxY, p[1])
	}

	return
}
