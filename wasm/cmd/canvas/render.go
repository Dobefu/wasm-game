package canvas

func Render() {
	for i := 0; i < 50; i++ {
		CONTEXT.Call("beginPath")
		CONTEXT.Call("moveTo", 0, 0)
		CONTEXT.Call("lineTo", 100, 100)
		CONTEXT.Call("stroke")
	}
}
