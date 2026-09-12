package bottlesong

import (
    "fmt"
    "strings"
)



var numStr  = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten"}

func Recite(startBottles, takeDown int) []string {
	s := make([]string, 0, 5*takeDown)

	for i := range takeDown {
		current := startBottles - i
		currentStr := numStr[current-1]
		switch current {
		case 10, 9, 8, 7, 6, 5, 4, 3:

			s = append(s, fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(currentStr)))
			s = append(s, fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(currentStr)))
			s = append(s, "And if one green bottle should accidentally fall,")
			s = append(s, fmt.Sprintf("There'll be %s green bottles hanging on the wall.", numStr[current-1-1]))
            if takeDown > 1 && i < takeDown - 1{
                s=append(s, "")
            }
            
		case 2:
			s = append(s, fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(currentStr)))
			s = append(s, fmt.Sprintf("%s green bottles hanging on the wall,", strings.Title(currentStr)))
			s = append(s, "And if one green bottle should accidentally fall,")
			s = append(s, fmt.Sprintf("There'll be %s green bottle hanging on the wall.", numStr[0]))
			if takeDown > 1 && i < takeDown - 1{
                s=append(s, "")
            }
        case 1:
			s = append(s, fmt.Sprintf("%s green bottle hanging on the wall,", strings.Title(currentStr)))
			s = append(s, fmt.Sprintf("%s green bottle hanging on the wall,", strings.Title(currentStr)))
			s = append(s, "And if one green bottle should accidentally fall,")
			s = append(s, "There'll be no green bottles hanging on the wall.")
			if takeDown > 1 && i < takeDown - 1{
                s=append(s, "")
            }
		}

	}
	return s
}
