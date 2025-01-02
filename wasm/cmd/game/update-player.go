package game

import (
	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

func UpdatePlayer(player *structs.GameObject) {
	if canMove(player, 0, DELTA_TIME*25) {
		player.Y += DELTA_TIME * 25
	}

	if canMove(player, DELTA_TIME*25, 0) {
		player.X += DELTA_TIME * 25
	}
}

func canMove(player *structs.GameObject, dx, dy float64) bool {
	tmpChar := &structs.GameObject{
		X:        player.X + dx,
		Y:        player.Y + dy,
		Rotation: player.Rotation,
		Scale:    player.Scale,
		Width:    player.Width,
		Height:   player.Height,
		Parent:   player.Parent,
	}

	for _, obj := range GAME_OBJECTS {
		if obj != player && CheckCollision(tmpChar, obj) {
			return false
		}
	}

	return true
}
