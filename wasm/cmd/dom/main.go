package dom

import "syscall/js"

var (
	DOCUMENT js.Value
)

func init() {
	DOCUMENT = js.Global().Get("document")
}
