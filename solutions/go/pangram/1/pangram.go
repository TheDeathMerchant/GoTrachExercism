package pangram

import (
    "strings"
    "slices"
    "regexp"
    )

func IsPangram(input string) bool {
	testStr := "abcdefghijklmnopqrstuvwxyz"
	input = strings.ToLower(input)
	m := make(map[rune]bool)
	for _, v := range input {
		m[v] = true
	}
	re := regexp.MustCompile(`[a-z]`)
	chars := make([]rune, 0, 26)
	for k := range m {
		if re.MatchString(string(k)) {
			chars = append(chars, k)
		}
	}

	slices.Sort(chars)
    return string(chars) == testStr
}
