package logs

// Application identifies the application emitting the given log.
func Application(log string) string {
	for _, val := range log {
        switch val {
            case '❗':
            	return "recommendation"
			case '🔍':
            	return "search"
            case '☀':
            	return "weather"
        }
    }
    return "default"
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	var runes []rune
    for _, val := range log {
        if val == oldRune {
            runes = append(runes, newRune)
        } else if val == newRune {
            runes = append(runes, oldRune)
        } else {
            runes = append(runes, val)
        }
    }
    return string(runes)
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	var count int
    for range log {
        count++
    }
    if count <= limit  {
        return true
    } else {
        return false
    }
}
