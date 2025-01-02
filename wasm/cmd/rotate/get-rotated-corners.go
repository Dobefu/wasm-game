package rotate

func GetRotatedCorners(cx, cy, width, height, rotation float64) [][2]float64 {
	halfW, halfH := width/2, height/2

	corners := [][2]float64{
		{-halfW, -halfH},
		{halfW, -halfH},
		{halfW, halfH},
		{-halfW, halfH},
	}

	rotatedCorners := make([][2]float64, 4)
	for i, corner := range corners {
		rotatedCorners[i][0], rotatedCorners[i][1] = RotatePoint(corner[0], corner[1], 0, 0, ToRadians(rotation))
		rotatedCorners[i][0] += cx
		rotatedCorners[i][1] += cy
	}

	return rotatedCorners
}
