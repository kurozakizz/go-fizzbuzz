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

func Test5ShouldSayBuzz(t *testing.T) {
	expected := "Buzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(5)
	assert(expected, actual, t)
}

func Test10ShouldSayBuzz(t *testing.T) {
	expected := "Buzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(10)
	assert(expected, actual, t)
}

func Test15ShouldSayFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(15)
	assert(expected, actual, t)
}

func Test30ShouldSayFizzBuzz(t *testing.T) {
	expected := "FizzBuzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(30)
	assert(expected, actual, t)
}

func assert(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
	}
}
