package structs

type GameObject struct {
	X, Y     float64
	Width    float64
	Height   float64
	Scale    float64
	Rotation float64
	Parent   *GameObject
	Type     GameObjectType
}
