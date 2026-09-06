package hamming

import "errors"

func Distance(a, b string) (int, error) {
    if len(a) != len(b) {
        return 0, errors.New("DNA strands are not of equal length")
    }
    count := 0
	for i, _ := range a{
        if a[i] != b[i] {
            count++
        } 
    }
    return count, nil
}
