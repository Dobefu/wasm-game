package game

import (
	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

func UpdatePlayer(player *structs.GameObject) {
	if !CheckCollision(*player) {
		player.Y += DELTA_TIME * 25
	}
}
