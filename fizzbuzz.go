package fizzbuzz

import "strconv"

const (
	fizzNumber  = 3
	buzzNumber  = 5
	fizzMessage = "Fizz"
	buzzMessage = "Buzz"
)

type FizzBuzz struct{}

func (f FizzBuzz) say(number int) string {
	isFizz := (number % fizzNumber) == 0
	isBuzz := (number % buzzNumber) == 0
	isFizzBuzz := isFizz && isBuzz
	if isFizzBuzz {
		return fizzMessage + buzzMessage
	}
	if isFizz {
		return fizzMessage
	}
	if isBuzz {
		return buzzMessage
	}
	return strconv.Itoa(number)
}
