package microblog

import "fmt"

func Truncate(phrase string) string {
	var s string
    for i, v := range []rune(phrase) {
        if i >= 5 {
            return s
        } else {
            s += fmt.Sprintf("%c", v)
        }
    }
    return s
}
