// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import (
    "strings"
)

func  engStrBuilder(s string) string{
    var b strings.Builder
    b.Grow(len(s))
    for _, v := range s {
        if (v >= 'a' && v <= 'z') || (v >='A' && v <= 'Z') || v == ' ' || v == '-' {
            if v == '-' {
                b.WriteRune(' ')
            } else {
                b.WriteRune(v)
            }
        }
    }
    return b.String()
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	words := strings.Fields(engStrBuilder(s))
    var out string
    for _, v := range words {
        out += strings.ToUpper(string(v[0]))
    }
	return out
}
