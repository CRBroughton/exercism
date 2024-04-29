package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, time int) int {
	if time != 0 {
		return len(layers) * time
	}
	return len(layers) * 2
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	noodles = 0
	sauce = 0.0

	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		}
		if layer == "sauce" {
			sauce += 0.2
		}
	}

	return noodles, sauce
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendsIngredients []string, myIngredients []string) {
	lastItemIndex := len(myIngredients) - 1
	myIngredients[lastItemIndex] = friendsIngredients[len(friendsIngredients)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(portions []float64, amount int) []float64 {
	var output []float64
	for _, portion := range portions {
		output = append(output, portion/2*float64(amount))
	}
	return output
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
