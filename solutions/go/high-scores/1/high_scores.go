package highscores

import (
    "slices"
    "fmt"
    )

type HighScores struct{
    scores []int
}

// NewHighScores returns a new HighScores object.
func NewHighScores(scores []int) *HighScores {
	return &HighScores{
        scores: scores,
    }
}

// Scores returns all the scores.
func (s *HighScores) Scores() []int {
	return s.scores
}

// Latest returns the latest (last) score.
func (s *HighScores) Latest() int {
	return s.scores[len(s.scores)-1]
}

// PersonalBest returns the best (highest) score.
func (s *HighScores) PersonalBest() int {
	scores := make([]int, len(s.scores))
    copy(scores, s.scores)
    slices.Sort(scores)
    slices.Reverse(scores)
    fmt.Println(scores[0])
    return scores[0]
    
}

// TopThree returns the top three scores.
func (s *HighScores) TopThree() []int {
   scores := make([]int, len(s.scores))
    copy(scores, s.scores)
    slices.Sort(scores)
    slices.Reverse(scores)
     if len(scores) < 3 {
        return scores
    }
    return scores[:3]
}
