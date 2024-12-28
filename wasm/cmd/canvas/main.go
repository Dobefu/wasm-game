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

	dom.AddEventListener("window", "resize", func() { resize(true) })
	resize(false)
}

func resize(rerender bool) {
	height := dom.WINDOW.Get("innerHeight").Int()
	width := dom.WINDOW.Get("innerWidth").Int()

	CANVAS.element.Set("height", height)
	CANVAS.element.Set("width", width)

	CANVAS.height = height
	CANVAS.width = width

	if rerender {
		Render()
	}
}
