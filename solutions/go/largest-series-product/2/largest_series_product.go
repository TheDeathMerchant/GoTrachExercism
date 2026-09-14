package largestseriesproduct

import (
    "errors"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("span must not exceed string length")
	} else if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	var largest int64
	//fmt.Println("len", len(products))
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for _, v := range digits[i : span+i] {
			if v < '0' || v > '9' {
				return 0, errors.New("digits input must only contain digits")
			}
			product *= int64(v - '0')
		}
		if product > largest {
            largest = product
        }
	}
	return largest, nil
}
