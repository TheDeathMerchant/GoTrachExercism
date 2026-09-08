package atbashcipher

import (
	"strings"
	"unicode"
)

func Atbash(s string) string {
	var b strings.Builder
	b.Grow(len(s) + len(s)/5) // Pre-allocate space to avoid reallocations

	count := 0 // Tracks letters/digits written to insert spaces every 5 chars

	for _, r := range s {
		r = unicode.ToLower(r)

		var out rune
		if r >= 'a' && r <= 'z' {
			out = 'a' + 'z' - r
		} else if r >= '0' && r <= '9' {
			out = r
		} else {
			continue // Skip spaces, punctuation, etc.
		}

		// Insert a space every 5 characters (except before the very first char)
		if count > 0 && count%5 == 0 {
			b.WriteByte(' ')
		}

		b.WriteRune(out)
		count++
	}

	return b.String()

}
