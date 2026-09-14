package largestseriesproduct

import (
    "errors"
    "slices"
)

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if len(digits) < span {
		return 0, errors.New("span must not exceed string length")
	} else if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	var products []int64
	products = make([]int64, len(digits)-span+1)
	//fmt.Println("len", len(products))
	for i := 0; i <= len(digits)-span; i++ {
		var product int64 = 1
		for _, v := range digits[i : span+i] {
			if v < '0' || v > '9' {
				return 0, errors.New("digits input must only contain digits")
			}
			product *= int64(v - '0')
		}
		products[i] = product
	}
	slices.Sort(products)
	return products[len(products)-1], nil
}
