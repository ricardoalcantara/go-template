package utils

type Number interface {
	uint | int | int32 | int64 | float32 | float64
}

func MapValue[T Number](x, inMin, inMax, outMin, outMax T) T {
	return (x-inMin)*(outMax-outMin)/(inMax-inMin) + outMin
}
