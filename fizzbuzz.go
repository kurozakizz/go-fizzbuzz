package fizzbuzz

import "strconv"

func say(number int) string {
	const fizzNumber = 3
	const buzzNumber = 5
	if number == fizzNumber {
		return "Fizz"
	}
	if number == buzzNumber {
		return "Buzz"
	}
	return strconv.Itoa(number)
}
