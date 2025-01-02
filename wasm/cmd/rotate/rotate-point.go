package rotate

import (
	"math"
)

func RotatePoint(x, y, cx, cy, angle float64) (float64, float64) {
	s, c := math.Sincos(angle)
	dx, dy := x-cx, y-cy

	return cx + (dx*c - dy*s), cy + (dx*s + dy*c)
}
