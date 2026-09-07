package luhn

import (
    "strings"
    
)

func Valid(id string) bool {
	
	id = strings.ReplaceAll(id, " ", "")
    if len(id) <= 1 {
		return false
	}
    for i  := range id {
        if id[i] < '0' || id[i] >'9' {
            return false
        }
    }
	var count int
	var j int
	numArr := make([]int, len(id))
	for i := len(id) - 1; i >= 0; i-- {
		vi := int(id[i] - '0')
		if j%2 != 0 {
			if vi*2 > 9 {
				numArr[i] = (vi * 2) - 9
			} else {
				numArr[i] = (vi * 2)
			}
		} else {
			numArr[i] = vi
		}
		j++
	}
	for _, v := range numArr {
		count += v
	}
	if count%10 == 0 {
		return true
	} else {
		return false
	}
}
