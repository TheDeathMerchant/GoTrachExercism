package scrabblescore

import "strings"

func Score(word string) int {
	var count int
    word = strings.ToLower(word)
    for _, v:= range word{
        switch v {
            case 'a', 'e', 'i', 'o', 'u', 'l', 'n', 'r', 's', 't':
            	count +=1
            case 'd', 'g':
            	count += 2
            case 'b', 'c', 'm', 'p':
            	count += 3
            case 'f', 'h', 'v', 'w', 'y':
            	count += 4
            case 'k':
            	count += 5
            case 'j', 'x':
            	count += 8
            case 'q', 'z':
            	count += 10
            default:
            	continue
        }
    }
    return count
}
