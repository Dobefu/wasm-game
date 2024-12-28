package canvas

func Render() {
	for i := 0; i < 50; i++ {
		CANVAS.context.Call("beginPath")
		CANVAS.context.Call("moveTo", 0, 0)
		CANVAS.context.Call("lineTo", CANVAS.width, CANVAS.height)
		CANVAS.context.Call("stroke")
	}
}
