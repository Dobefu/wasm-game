package dom

import "syscall/js"

var (
	DOCUMENT js.Value
	WINDOW   js.Value
)

func init() {
	DOCUMENT = js.Global().Get("document")
	DOCUMENT = js.Global().Get("window")
}
