package game

import "github.com/Dobefu/wasm-game/cmd/game/structs"

func Instantiate(gameObject structs.GameObject) {
	GAME_OBJECTS = append(GAME_OBJECTS, gameObject)
}