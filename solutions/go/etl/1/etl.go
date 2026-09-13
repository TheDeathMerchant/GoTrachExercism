package etl

import ("strings")

func Transform(in map[int][]string) map[string]int {
	scores := make(map[string]int)

    for k, v := range in {
        for _, letter := range v {
            scores[strings.ToLower(letter)] = k
        }
    }
    return scores
}
