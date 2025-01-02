package game

func Render(clearCanvas bool) {
	if clearCanvas {
		CANVAS.Context.Set("fillStyle", "black")
		CANVAS.Context.Call("fillRect", 0, 0, CANVAS.Width, CANVAS.Height)
	}

	for _, gameObject := range GAME_OBJECTS {
		RenderGameObject(gameObject)
	}
}
