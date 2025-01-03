//go:build js && wasm
// +build js,wasm

package controls

import (
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/controls/structs"
	"github.com/Dobefu/wasm-game/cmd/dom"
)

var Keys structs.Keys

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
		Keys.Left = true
		break
	case key == "ArrowRight" || key == "l" || key == "d":
		Keys.Right = true
		break
	case key == "ArrowUp" || key == "k" || key == "w" || key == " ":
		Keys.Up = true
		break
	case key == "ArrowDown" || key == "j" || key == "s":
		Keys.Down = true
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
		Keys.Left = false
		break
	case key == "ArrowRight" || key == "l" || key == "d":
		Keys.Right = false
		break
	case key == "ArrowUp" || key == "k" || key == "w" || key == " ":
		Keys.Up = false
		break
	case key == "ArrowDown" || key == "j" || key == "s":
		Keys.Down = false
		break
	}
}
