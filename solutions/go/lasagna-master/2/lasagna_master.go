package lasagnamaster
// TODO: define the 'PreparationTime()' function

func PreparationTime(layers []string, avgPreparationTime int) int {
    if avgPreparationTime == 0 {
        return len(layers) * 2
    }
    return  len(layers) * avgPreparationTime 
}
// TODO: define the 'Quantities()' function

func Quantities(layers []string) (noodlesAmount int, sauceAmount float64) {

    for _, layer := range layers {
        if layer == "noodles" {
            noodlesAmount += 50
        } else if layer == "sauce" {
			sauceAmount += 0.2
        }
    }
    return noodlesAmount, sauceAmount
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(fLayers, mLayers []string) {
    mLayers[len(mLayers)-1] = fLayers[len(fLayers)-1]
}
// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(q1 []float64, portion int) (q2 []float64) {
    q2 = make([]float64, len(q1))
    copy(q2, q1)
    req := float64(portion)/2.0
    for i := range q2 {
        q2[i] *= req
    }
	return q2
}
// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
