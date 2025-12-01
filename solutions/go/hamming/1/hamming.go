package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, errors.New("strings must be of the same length")
	}
	distance := 0
	for i,v := range a {
		if v != rune(b[i]) {
			distance+=1
		}
	}
	return distance, nil
}
