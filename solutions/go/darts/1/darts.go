package darts

import "math"

func radius(x, y float64) float64 {
    return math.Sqrt(x*x + y*y)
}
func Score(x, y float64) int {
	r := radius(x, y)

    switch {
        case r <= 1:
        	return 10
        case r <= 5:
        	return 5
        case r <= 10:
        	return 1
        default:
        	return 0
    }
}
