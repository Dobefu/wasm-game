package dom

import (
	"syscall/js"
)

func wrapGoFunction(fn func(js.Value, []js.Value)) func(js.Value, []js.Value) interface{} {
	return func(this js.Value, args []js.Value) interface{} {
		fn(this, args)
		return nil
	}
}

func AddEventListener(element string, event string, fn func(js.Value, []js.Value)) {
	switch element {
	case "document":
		DOCUMENT.Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
		break
	case "window":
		WINDOW.Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
		break
	default:
		QuerySelector(element).Call("addEventListener", event, js.FuncOf(wrapGoFunction(fn)))
		break
	}
}
