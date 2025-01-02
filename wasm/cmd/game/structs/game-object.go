package structs

type GameObject struct {
	X        float64
	Y        float64
	Width    float64
	Height   float64
	Rotation float64
	Parent   *GameObject
}
