package game

import "math"

//Round to the nearest integer
func Round(f float64) float64 {
	return math.Floor(f + .5)
}
