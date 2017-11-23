package fizzbuzz

import (
	"testing"
)

type testData struct {
	number   int
	expected string
}

func TestNonFizzBuzzNumberShouldSayNonFizzBuzzNumber(t *testing.T) {
	testList := []testData{
		testData{number: 1, expected: "1"},
		testData{number: 2, expected: "2"},
		testData{number: 4, expected: "4"},
		testData{number: 7, expected: "7"},
		testData{number: 8, expected: "8"},
		testData{number: 11, expected: "11"},
		testData{number: 13, expected: "13"},
		testData{number: 14, expected: "14"},
	}
	fizzBuzz := FizzBuzz{}

	for _, data := range testList {
		actual := fizzBuzz.say(data.number)
		assert(data.expected, actual, t)
	}
}

func TestFizzNumberShouldSayFizz(t *testing.T) {
	testList := []testData{
		testData{number: 3, expected: fizzMessage},
		testData{number: 6, expected: fizzMessage},
		testData{number: 9, expected: fizzMessage},
		testData{number: 12, expected: fizzMessage},
		testData{number: 99, expected: fizzMessage},
	}
	fizzBuzz := FizzBuzz{}

	for _, data := range testList {
		actual := fizzBuzz.say(data.number)
		assert(data.expected, actual, t)
	}
}

func TestBuzzNumberShouldSayBuzz(t *testing.T) {
	testList := []testData{
		testData{number: 5, expected: buzzMessage},
		testData{number: 10, expected: buzzMessage},
		testData{number: 20, expected: buzzMessage},
		testData{number: 25, expected: buzzMessage},
		testData{number: 35, expected: buzzMessage},
	}
	fizzBuzz := FizzBuzz{}

	for _, data := range testList {
		actual := fizzBuzz.say(data.number)
		assert(data.expected, actual, t)
	}
}

func TestFizzBuzzNumberShouldSayFizzBuzz(t *testing.T) {
	expected := fizzMessage + buzzMessage
	testList := []testData{
		testData{number: 15, expected: expected},
		testData{number: 30, expected: expected},
		testData{number: 60, expected: expected},
		testData{number: 90, expected: expected},
	}
	fizzBuzz := FizzBuzz{}

	for _, data := range testList {
		actual := fizzBuzz.say(data.number)
		assert(data.expected, actual, t)
	}
}

func assert(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
	}
}
