package game

import "github.com/Dobefu/wasm-game/cmd/game/structs"

func UpdatePlayer(player *structs.GameObject) {
	player.Y += DELTA_TIME * 100
}
