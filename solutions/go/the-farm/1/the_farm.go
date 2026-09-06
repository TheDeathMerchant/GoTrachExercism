package thefarm
import (
    "fmt"
    "errors"
)
// TODO: define the 'DivideFood' function
func DivideFood(fc FodderCalculator, n int) (float64, error) {
    fa, err := fc.FodderAmount(n)
    if err != nil {
        return 0.0, err
    }

    ff, err := fc.FatteningFactor()

    if err != nil {
        return 0.0, err
    }

    return ff*fa/float64(n), nil
}
// TODO: define the 'ValidateInputAndDivideFood' function
func ValidateInputAndDivideFood(fc FodderCalculator, n int) (float64, error) {
    if n <= 0 {
        return 0.0, errors.New("invalid number of cows")
    } 
    return DivideFood(fc, n)
}
// TODO: define the 'ValidateNumberOfCows' function
type InvalidCowsError struct{
    n int
    errMsg string
}

func (i *InvalidCowsError) Error() string {
    return fmt.Sprintf("%d cows are invalid: %s", i.n, i.errMsg)
}

func ValidateNumberOfCows(n int) error {
    if n < 0 {
        return &InvalidCowsError{
            n: n,
            errMsg: "there are no negative cows",
            }
        } else if n == 0 {
            return &InvalidCowsError{
                n: n,
                errMsg: "no cows don't need food",
            }
        } else {
        	return nil
        }    
    }

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
