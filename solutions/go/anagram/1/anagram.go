package anagram

import (
    "strings"
    "slices"
    )

func Detect(subject string, candidates []string) []string {
	subject = strings.ToLower(subject)
    subjectArr := []rune(subject)
    slices.Sort(subjectArr)
	outArr := make([]string, 0, len(candidates))
    for _, v:=range candidates {
        vLower := strings.ToLower(v)
        if vLower == subject {
            continue
        }
        vArr := []rune(vLower)
        slices.Sort(vArr)
        if slices.Equal(vArr, subjectArr) {
            outArr = append(outArr, v)
        }
    }
    return outArr
}
