package game

import (
	"math"
	"time"

	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

func Update() {
	FRAME = FRAME + 1

	DELTA_TIME = math.Min(float64(time.Now().UnixMilli()-_LAST_TIME.UnixMilli())/1000, .1)
	_LAST_TIME = time.Now()

	for _, gameObject := range GAME_OBJECTS {
		if gameObject.Type == structs.GAME_OBJECT_TYPE_PLAYER {
			UpdatePlayer(gameObject)
		}
	}
}
