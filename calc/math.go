package calc

import "errors"

// returns sum of two integers
func Add(numbers ...int) int {
	sum := 0

	if len(numbers) < 2 {
		return errors.New("provude more than 2 numbers"), sum
	} else {
		for _, num := range numbers {
			sum = sum + num
		}

		return nil, sum
	}
}
