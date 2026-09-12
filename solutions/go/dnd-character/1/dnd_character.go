package dndcharacter

import (
    "math/rand"
    "slices"
    )

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
    mod := (score - 10)
	if mod < 0 && mod % 2 != 0 {
		return mod/ 2 - 1
    }
    return mod/2
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
    dieScores := [4]int{}
	for i:= range 4 {
        dieScores[i] = rand.Intn(6) + 1
    }
    slices.Sort(dieScores[:])
    sum := 0
    for i := range 3 {
        sum += dieScores[i] 
    }
    return sum
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	c1 := Character {
        Strength     : Ability(),
    	Dexterity    : Ability(),
    	Constitution : Ability(),
    	Intelligence : Ability(),
    	Wisdom       : Ability(),
    	Charisma     : Ability(),
    	Hitpoints    : 10,
    }
    c1.Hitpoints += Modifier(c1.Constitution)
    return c1
}
