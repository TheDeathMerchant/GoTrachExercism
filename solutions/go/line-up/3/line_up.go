package lineup

import "fmt"

func Format(name string, number int) string {
	var ordNum = "th"
     
    if number % 10 == 1 && number%100 != 11 {
        ordNum = "st"
    } else if number % 10 == 2 && number%100 != 12 {
        ordNum = "nd"
    } else if number % 10 == 3 && number%100 != 13 {
        ordNum = "rd"
    } else {
        ordNum = "th"
    }
    msg :=  fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, ordNum)
    return msg
}
