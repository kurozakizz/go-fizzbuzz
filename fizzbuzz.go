package fizzbuzz

import "strconv"

type FizzBuzz struct {
	fizzNumber  int
	buzzNumber  int
	fizzMessage string
	buzzMessage string
}

func newFizzBuzz() *FizzBuzz {
	f := FizzBuzz{
		fizzNumber:  3,
		buzzNumber:  5,
		fizzMessage: "Fizz",
		buzzMessage: "Buzz",
	}
	return &f
}

func (f FizzBuzz) say(number int) string {
	isFizz := (number % f.fizzNumber) == 0
	isBuzz := (number % f.buzzNumber) == 0
	if isFizz {
		return f.fizzMessage
	}
	if isBuzz {
		return f.buzzMessage
	}
	return strconv.Itoa(number)
}
