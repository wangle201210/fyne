package scale

import (
	"math"

	"fyne.io/fyne/v2"
)

// ToScreenCoordinate converts a fyne coordinate in the given canvas to a screen coordinate
func ToScreenCoordinate(c fyne.Canvas, v float32) int {
	return int(math.Ceil(float64(v * c.Scale())))
}

// ToFyneCoordinate converts a screen coordinate for a given canvas to a fyne coordinate
func ToFyneCoordinate(c fyne.Canvas, v int) float32 {
	switch c.Scale() {
	case 0.0:
		panic("Incorrect scale most likely not set.")
	case 1.0:
		return float32(v)
	default:
		return float32(v) / c.Scale()
	}
}
