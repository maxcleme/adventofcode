package utils

import "math"

func RequiredFuel(mass float64) float64 {
	return math.Floor(mass/3.0) - 2
}
