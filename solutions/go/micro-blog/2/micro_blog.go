package microblog

func Truncate(phrase string) string {
    runes := []rune(phrase)
    if len(phrase) <5 {
        return phrase
    }
    return string(runes[:5])
}
