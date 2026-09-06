package parsinglogfiles

import (
    "regexp"
    "fmt"
    )
var (
    validLineRe = regexp.MustCompile(`^\[(TRC|DBG|INF|WRN|ERR|FTL)\]`)
    splitLineRe = regexp.MustCompile(`<[~*=-]*>`)
    quotedPassRe = regexp.MustCompile(`(?i)".*?password.*?"`)
    removeLineRe = regexp.MustCompile(`end-of-line[0-9]+`)
    tagUserRe = regexp.MustCompile(`User\s+(\w+)`)
)
func IsValidLine(text string) bool {
	return validLineRe.MatchString(text)
}

func SplitLogLine(text string) []string {
    return splitLineRe.Split(text, -1)
}

func CountQuotedPasswords(lines []string) int {
	var count int
	for _, line := range lines {
		if quotedPassRe.MatchString(line) {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	return removeLineRe.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
    var s []string
    for _, line := range lines {
        match := tagUserRe.FindStringSubmatch(line)
        if len(match) > 1 {
            str := fmt.Sprintf("[USR] %s %s", match[1], line)
            s = append(s, str)
        } else {
            s = append(s, line)
        }
    }
    return s
}
