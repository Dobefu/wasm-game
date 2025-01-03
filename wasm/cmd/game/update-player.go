//go:build js && wasm
// +build js,wasm

package game

import (
	"fmt"
	"math"

	"github.com/Dobefu/wasm-game/cmd/controls"
	"github.com/Dobefu/wasm-game/cmd/game/structs"
)

func UpdatePlayer(player *structs.GameObject) {
	fmt.Println(player.XSpeed)
	if canMove(player, player.XSpeed, -.01) {
		player.X += player.XSpeed
	} else {
		for canMove(player, DELTA_TIME*25, -.01) {
			player.X += DELTA_TIME * 25
		}

		player.X = math.Round(player.X)
		player.XSpeed = 0
	}

	if canMove(player, 0, player.YSpeed) {
		player.YSpeed += DELTA_TIME * 25
		player.Y += player.YSpeed
	} else {
		for canMove(player, 0, DELTA_TIME*25) {
			player.Y += DELTA_TIME * 25
		}

		player.Y = math.Round(player.Y)
		player.YSpeed = 0
	}

	var xsp float64

	if controls.Keys.Left {
		xsp = -DELTA_TIME * 4
	}

	if controls.Keys.Right {
		xsp = DELTA_TIME * 4
	}

	player.XSpeed = math.Max(math.Min(player.XSpeed+xsp, 1), -1)
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
