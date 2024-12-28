package dom

import "syscall/js"

func RequestAnimationFrame(fn func()) {
	WINDOW.Call("requestAnimationFrame", js.FuncOf(func(_ js.Value, _ []js.Value) interface{} {
		RequestAnimationFrame(fn)
		return nil
	}))
	fn()
}
