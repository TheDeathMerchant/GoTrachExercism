package grains

import (
    "math"
    "errors"
)

func Square(number int) (uint64, error) {
	if number > 64 {
        return 0, errors.New("square greater than 64 is invalid")
    } else if number < 0 {
        return 0, errors.New("negative square is invalid")
    } else if number == 0 {
        return 0, errors.New("square 0 is invalid")
    }
    return uint64(math.Pow(2, float64(number-1))), nil
}

func Total() uint64 {
	var sum uint64
    for i := range 64{
        sum += uint64(math.Pow(2, float64(i)))
    }
    return sum
}
