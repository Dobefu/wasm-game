package canvas

import (
	"github.com/Dobefu/wasm-game/cmd/dom"
)

var (
	CANVAS Canvas
)

func init() {
	CANVAS = Canvas{element: dom.GetElementById("game")}
	CANVAS.context = CANVAS.element.Call("getContext", "2d")

	dom.AddEventListener("window", "resize", func() { resizeCanvas(true) })
	resizeCanvas(false)
}

func resizeCanvas(rerender bool) {
	CANVAS.element.Set("height", dom.WINDOW.Get("innerHeight"))
	CANVAS.element.Set("width", dom.WINDOW.Get("innerWidth"))

	if rerender {
		Render()
	}
}
