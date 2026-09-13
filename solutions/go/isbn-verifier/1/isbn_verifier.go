package isbnverifier

import (
    "strings"
	"strconv"
    )

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")
    sum := 0
    if len(isbn) != 10 {
            return false
        }
    for i, v := range isbn {
        
        if v >= '0' && v <= '9' {
            num, _ := strconv.Atoi(string(v))
            sum += num * (10-i)
        } else if v == 'X' && i == len(isbn)-1 {
            sum += 10 
        } else {
            return false
        }
    }
    if sum % 11 == 0 {
        return true
    } else {
        return false
    }
}
