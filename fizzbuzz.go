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
	if number == f.fizzNumber {
		return f.fizzMessage
	}
	if number == f.buzzNumber {
		return f.buzzMessage
	}
	return strconv.Itoa(number)
}
