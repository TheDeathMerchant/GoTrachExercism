// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
    "strings"
    "unicode"
    "fmt"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
     remark = strings.TrimSpace(remark)
    fmt.Println(remark)
	isQues := false
	isYell := false
	allSpace := true
	if strings.HasSuffix(remark, "?") {
		isQues = true
	}
	outer:
	for _, v := range remark {
        if !unicode.IsSpace(v) {
			allSpace = false
			}
		if unicode.IsLetter(v) {
 
            if unicode.IsUpper(v) {
				isYell = true
			} else if !unicode.IsUpper(v) {
				isYell = false
				break outer
			}
		} 
        fmt.Printf("%c\n", v)
	}

	switch {
	case isQues && isYell:
		return "Calm down, I know what I'm doing!"
	case isQues && !isYell:
		return "Sure."
	case isYell && !isQues:
		return "Whoa, chill out!"
	case allSpace:
		return "Fine. Be that way!"
	default:
		return "Whatever."
	}
}
