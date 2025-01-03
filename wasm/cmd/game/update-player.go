//go:build js && wasm
// +build js,wasm

package game

import (
	"math"

	"github.com/Dobefu/wasm-game/cmd/controls"
	"github.com/Dobefu/wasm-game/cmd/game/structs"
	"github.com/Dobefu/wasm-game/cmd/utils"
)

func UpdatePlayer(player *structs.GameObject) {
	if canMove(player, player.XSpeed, 0) {
		player.X += player.XSpeed
	} else {
		for canMove(player, DELTA_TIME*25*float64(utils.Sign(player.XSpeed)), 0) {
			player.X += DELTA_TIME * 25 * float64(utils.Sign(player.XSpeed))
		}

		player.X = math.Round(player.X)
		player.X += float64(utils.Sign(player.XSpeed)) * -1e-6
		player.XSpeed = 0
	}

	if canMove(player, 0, player.YSpeed+DELTA_TIME*25) {
		player.YSpeed += DELTA_TIME * 25
		player.Y += player.YSpeed
	} else {
		for canMove(player, 0, DELTA_TIME*25) {
			player.Y += DELTA_TIME * 25
		}

		player.Y = math.Round(player.Y)
		player.Y -= 1e-6
		player.YSpeed = 0
	}

	var xsp float64

	if controls.Keys.Left {
		xsp = -DELTA_TIME * 4
	}

	if controls.Keys.Right {
		xsp = DELTA_TIME * 4
	}

	player.XSpeed = math.Max(math.Min(player.XSpeed+xsp, 4), -4)

	if xsp == 0 {
		player.XSpeed *= .9

		if math.Abs(player.XSpeed) <= .01 {
			player.XSpeed = 0
		}
	}

	if !canMove(player, 0, 1e-6) && controls.Keys.Up {
		player.YSpeed = -DELTA_TIME * 500
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
