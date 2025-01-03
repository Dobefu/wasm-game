package canvas

import (
	"syscall/js"

	"github.com/Dobefu/wasm-game/cmd/dom"
)

var (
	CANVAS Canvas
)

func init() {
	CANVAS = Canvas{Element: dom.GetElementById("game")}
	CANVAS.Context = CANVAS.Element.Call("getContext", "2d")

	dom.AddEventListener("window", "resize", func(_ js.Value, _ []js.Value) { resize() })
	resize()
}

func resize() {
	height := dom.WINDOW.Get("innerHeight").Int()
	width := dom.WINDOW.Get("innerWidth").Int()

	CANVAS.Element.Set("height", height)
	CANVAS.Element.Set("width", width)

	CANVAS.Height = height
	CANVAS.Width = width
}
