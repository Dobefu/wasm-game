package dom

import "syscall/js"

func QuerySelector(query string) js.Value {
	return DOCUMENT.Call("querySelector", query)
}
