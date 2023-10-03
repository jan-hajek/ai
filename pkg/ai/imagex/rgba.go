package imagex

// RgbaToPixel img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func RgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct example
type Pixel struct {
	R int
	G int
	B int
	A int
}
