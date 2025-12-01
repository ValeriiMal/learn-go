package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, errors.New("n must be positive")
	}
	steps := 0
	if n == 1 {
		return steps, nil
	}
	for n > 1 {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = n * 3 + 1
		}
		steps++
	}
	return steps, nil
}
