package fizzbuzz

import "strconv"

func say(number int) string {
	const fizzNumber = 3
	if number == fizzNumber {
		return "Fizz"
	}
	return strconv.Itoa(number)
}
