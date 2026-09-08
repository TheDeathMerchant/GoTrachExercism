package armstrongnumbers

import "math"

func IsNumber(n int) bool {
	var sum int
    if n == 0{
        return true
    }
    nCopy := n
	nums := []int{}
    for n > 0 {
        nums = append(nums, n%10)
        n = n/10
    }
    for _, i := range nums {
        sum += int(math.Pow(float64(i), float64(len(nums))))
    }
    return sum == nCopy 
}
