package lasagna

func PreparationTime(layers []string, preparationTime int) int {
	if preparationTime == 0 {
		preparationTime = 2
	}
	return len(layers) * preparationTime
}

func Quantities(parts []string) (int, float64) {
	sauceCount := 0
	for _, v := range parts {
		if v == "sauce" {
			sauceCount++
		}
	}
	noodlesCount := 0
	for _, v := range parts {
		if v == "noodles" {
			noodlesCount++
		}
	}
	return noodlesCount * 50, float64(sauceCount) * 0.2
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
	amountsForOnePerson := []float64{}
	for _, v := range amounts {
		amountsForOnePerson = append(amountsForOnePerson, v*float64(portions)/2)
	}
	return amountsForOnePerson
}
