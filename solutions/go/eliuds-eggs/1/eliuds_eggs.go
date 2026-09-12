package eliudseggs

import "strconv"

func EggCount(displayValue int) int {
	num := strconv.FormatInt(int64(displayValue), 2)
    numBin, _ := strconv.ParseInt(num, 2, 64)
    count := 0
    for numBin > 0 {
        numBin &= numBin-1
        count++
    }
    return count
}
