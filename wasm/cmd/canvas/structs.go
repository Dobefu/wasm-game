package canvas

import "syscall/js"

type Canvas struct {
	height  int
	width   int
	element js.Value
	context js.Value
}
