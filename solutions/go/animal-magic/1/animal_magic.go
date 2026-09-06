package chance

import "math/rand"
// RollADie returns a random int d with 1 <= d <= 20.
func RollADie() int {
	var r int
    for r = rand.Intn(21);r <1 || r>20;r = rand.Intn(21) {}
    return r
}

// GenerateWandEnergy returns a random float64 f with 0.0 <= f < 12.0.
func GenerateWandEnergy() float64 {
	return rand.Float64() * float64(12.0)
}

// ShuffleAnimals returns a slice with all eight animal strings in random order.
func ShuffleAnimals() []string {
	animals := []string{"ant", "beaver", "cat", "dog", "elephant", "fox", "giraffe", "hedgehog"}
    rand.Shuffle(len(animals), func (i, j int) {
        animals[i], animals[j]  =  animals[j], animals[i]
    }) 
    return animals
    
}
