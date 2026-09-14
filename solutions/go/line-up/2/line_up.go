package lineup

import "fmt"

func Format(name string, number int) string {
	var ordNum string
    

    if number > 100 && number%100 == 11 || number%100 == 12 || number%100 == 13 {
        ordNum = "th"
        msg :=  fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, ordNum)
        return msg
    } 
    if number % 10 == 1 && number != 11 {
        ordNum = "st"
    } else if number % 10 ==  2 {
        ordNum = "nd"
    } else if number % 10 == 3 {
        ordNum = "rd"
    } else {
        ordNum = "th"
    }
    msg :=  fmt.Sprintf("%s, you are the %d%s customer we serve today. Thank you!", name, number, ordNum)
    return msg
}
