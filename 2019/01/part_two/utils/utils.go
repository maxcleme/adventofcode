package utils

import "math"

func RequiredFuel(mass float64) float64 {
	fuel := math.Floor(mass/3.0) - 2
	// exit case
	if fuel < 0 {
		return 0
	}
	return fuel + RequiredFuel(fuel)
}
