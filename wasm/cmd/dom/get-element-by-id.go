package dom

import "syscall/js"

func GetElementById(element string) js.Value {
	return DOCUMENT.Call("getElementById", element)
}
