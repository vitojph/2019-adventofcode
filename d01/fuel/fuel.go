package fuel

import (
	"math"
)

// Fuel computes the fuel needed to launch a module based on its mass
func Fuel(mass int) int {
	fmass := float64(mass)
	fuel := int((math.Floor(fmass / 3)) - 2)
	if fuel > 0 {
		return fuel + Fuel(fuel)
	}
	return 0
}
