package game

import (
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

	obj := Instantiate(structs.GameObject{Width: 100, Height: 100, X: 500, Y: 500, Scale: .5})
	Instantiate(structs.GameObject{Width: 100, Height: 100, X: 150, Parent: obj})
	Instantiate(structs.GameObject{Width: 100, Height: 100, Y: 150, Parent: obj, Scale: .5})
}
