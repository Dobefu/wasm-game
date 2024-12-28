package canvas

import (
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/dom"
)

var (
	CANVAS  js.Value
	CONTEXT js.Value
)

func init() {
	CANVAS = dom.GetElementById("game")
	CONTEXT = CANVAS.Call("getContext", "2d")

	dom.AddEventListener("window", "resize", func() { resizeCanvas(true) })
	resizeCanvas(false)
}

func resizeCanvas(rerender bool) {
	CANVAS.Set("height", dom.WINDOW.Get("innerHeight"))
	CANVAS.Set("width", dom.WINDOW.Get("innerWidth"))

	if rerender {
		Render()
	}
}
