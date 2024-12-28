package dom

import "syscall/js"

func wrapGoFunction(fn func()) func(js.Value, []js.Value) interface{} {
	return func(_ js.Value, _ []js.Value) interface{} {
		fn()
		return nil
	}
}

func AddEventListener(element string, event string, fn func()) {
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
