package collatzconjecture
import (
    
    "errors"
)
func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
        return 0, errors.New("n is not a positive integer")
    }
    
    var i int
    for i = 0; n!=1; i++ {
        if n%2 == 0{
            n = n/2
        } else if n%2 != 0 {
            n = n*3 + 1
        }
    }
    return i, nil
}
