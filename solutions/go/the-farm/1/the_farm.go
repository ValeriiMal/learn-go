package thefarm

import (
	"errors"
	"fmt"
)

func DivideFood(calc FodderCalculator, number int) (float64, error) {
	amount, amountErr := calc.FodderAmount(number)
	if amountErr != nil {
		return 0, amountErr
	}
	fatt,fattErr := calc.FatteningFactor()
	if fattErr != nil {
		return 0, fattErr
	}
	return (amount * fatt) / float64(number), nil
}

func ValidateInputAndDivideFood(calc FodderCalculator, number int) (float64, error) {
	if number > 0 {
		return DivideFood(calc, number)
	}
	return 0, errors.New("invalid number of cows")
}

type InvalidCowsError struct {
	number int
	message string
}

func (err *InvalidCowsError) Error() string {
	return fmt.Sprintf("%d cows are invalid: %s", err.number, err.message)
}

func ValidateNumberOfCows(number int) error {
	if number < 0 {
		return &InvalidCowsError{
			number: number,
			message: "there are no negative cows",
		}
	}
	if number == 0 {
		return &InvalidCowsError{
			number: number,
			message: "no cows don't need food",
		}
	}
	return nil
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
