package game

import (
	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

func Instantiate(gameObject structs.GameObject) *structs.GameObject {
	if gameObject.Scale == 0 {
		gameObject.Scale = 1
	}

	GAME_OBJECTS = append(GAME_OBJECTS, &gameObject)

	return &gameObject
}
