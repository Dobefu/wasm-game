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

	Instantiate(structs.GameObject{Width: 10000, Height: 10, X: 0, Y: 800})
	Instantiate(structs.GameObject{Width: 50, Height: 50, X: 200, Y: 770})
	Instantiate(structs.GameObject{Width: 50, Height: 50, X: 250, Y: 770})
	Instantiate(structs.GameObject{Width: 50, Height: 50, X: 450, Y: 770})
	Instantiate(structs.GameObject{Width: 50, Height: 50, X: 550, Y: 770})
	Instantiate(structs.GameObject{Width: 50, Height: 50, X: 350, Y: 720})
}
