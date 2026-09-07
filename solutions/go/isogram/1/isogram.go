package isogram

import "strings"

func IsIsogram(word string) bool {
    word = strings.ToLower(word)
	m := make(map[string]int)
    for _, l := range word {
        m[string(l)]++
    }
    for k, v := range m {
        if v > 1 && k!=" " && k!="-" {
            return false
        }
    }
    return true
}
