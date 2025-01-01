package structs

type GameObject struct {
	X        float64
	Y        float64
	Rotation float64
	Parent   *GameObject
}
