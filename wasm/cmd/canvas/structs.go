package canvas

import "syscall/js"

type Canvas struct {
	height  uint8
	width   uint8
	element js.Value
	context js.Value
}
