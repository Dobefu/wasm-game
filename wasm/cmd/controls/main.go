//go:build js && wasm
// +build js,wasm

package controls

import (
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/controls/structs"
	"github.com/Dobefu/wasm-game/cmd/dom"
)

var state structs.Keys

func init() {
	dom.AddEventListener("window", "keydown", handleKeyDown)
	dom.AddEventListener("window", "keyup", handleKeyUp)
}

func handleKeyDown(_ js.Value, args []js.Value) {
	if len(args) != 1 {
		return
	}

	key := args[0].Get("key").String()

	switch {
	case key == "ArrowLeft" || key == "h" || key == "a":
		state.Left = true
		break
	case key == "ArrowRight" || key == "l" || key == "d":
		state.Right = true
		break
	case key == "ArrowUp" || key == "k" || key == "w" || key == " ":
		state.Up = true
		break
	case key == "ArrowDown" || key == "j" || key == "s":
		state.Down = true
		break
	}
}

func handleKeyUp(_ js.Value, args []js.Value) {
	if len(args) != 1 {
		return
	}

	key := args[0].Get("key").String()

	switch {
	case key == "ArrowLeft" || key == "h" || key == "a":
		state.Left = false
		break
	case key == "ArrowRight" || key == "l" || key == "d":
		state.Right = false
		break
	case key == "ArrowUp" || key == "k" || key == "w" || key == " ":
		state.Up = false
		break
	case key == "ArrowDown" || key == "j" || key == "s":
		state.Down = false
		break
	}
}
