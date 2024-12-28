package canvas

func Render() {
	CANVAS.context.Call("beginPath")
	CANVAS.context.Call("moveTo", 0, 0)
	CANVAS.context.Call("lineTo", CANVAS.width, CANVAS.height)
	CANVAS.context.Call("stroke")

	CANVAS.context.Call("beginPath")
	CANVAS.context.Call("moveTo", CANVAS.width, 0)
	CANVAS.context.Call("lineTo", 0, CANVAS.height)
	CANVAS.context.Call("stroke")
}
