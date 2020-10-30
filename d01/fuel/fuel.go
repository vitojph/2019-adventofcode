package fuel

import (
	"math"
)

// Fuel computes the fuel needed to launch a module based on its mass
func Fuel(mass int) int {
	fmass := float64(mass)
	ffuel := (math.Floor(fmass / 3)) - 2
	return int(ffuel)
}
