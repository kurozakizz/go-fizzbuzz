package fizzbuzz

import "strconv"

func say(number int) string {
	const fizzNumber = 3
	const buzzNumber = 5
	const fizzMessage = "Fizz"
	const buzzMessage = "Buzz"
	if number == fizzNumber {
		return fizzMessage
	}
	if number == buzzNumber {
		return buzzMessage
	}
	return strconv.Itoa(number)
}
