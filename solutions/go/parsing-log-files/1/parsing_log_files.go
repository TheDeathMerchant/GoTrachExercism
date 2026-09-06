package parsinglogfiles

import (
    "regexp"
    "fmt"
    )

func IsValidLine(text string) bool {
	var result bool
	logs := []string{"TRC", "DBG", "INF", "WRN", "ERR", "FTL"}
	for _, log := range logs {
		re, err := regexp.Compile(fmt.Sprintf(`^\[%s\]`, log))
		if err != nil {
			fmt.Println(err)
		}
		result = re.MatchString(text)
		if result {
			break
		}
	}
	return result
}

func SplitLogLine(text string) []string {
	re := regexp.MustCompile(`<[~*=-]*>`)
	result := re.Split(text, -1)
    return result
}

func CountQuotedPasswords(lines []string) int {
	re := regexp.MustCompile(`(?i)".*?password.*?"`)
	var count int
	for _, line := range lines {
		result := re.MatchString(line)
		//fmt.Println(line, result)
		if result {
			count++
		}
	}
	return count
}

func RemoveEndOfLineText(text string) string {
	re := regexp.MustCompile(`end-of-line[0-9]+`)
	return re.ReplaceAllString(text, "")

}

func TagWithUserName(lines []string) []string {
	re := regexp.MustCompile(`User\s+(\w+)`)
    var s []string
    for _, line := range lines {
        match := re.FindStringSubmatch(line)
        if len(match) > 1 {
            str := fmt.Sprintf("[USR] %s %s", match[1], line)
            s = append(s, str)
        } else {
            s = append(s, line)
        }
    }
    return s
}
