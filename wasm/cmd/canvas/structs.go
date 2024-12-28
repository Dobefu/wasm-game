package canvas

import "syscall/js"

type Canvas struct {
	Height  int
	Width   int
	Element js.Value
	Context js.Value
}
