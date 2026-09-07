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
	for i := len(id) - 1; i >= 0; i-- {
		vi := int(id[i] - '0')
		if j%2 != 0 {
			if vi*2 > 9 {
                count += (vi * 2) - 9
			} else {
                count += (vi * 2)
			}
		} else {
            count += vi
		}
		j++
    }
	if count%10 == 0 {
		return true
	} else {
		return false
	}
}
