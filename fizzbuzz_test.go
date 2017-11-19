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

func Test1ShouldSay1(t *testing.T) {
	expected := "1"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(1)
	assert(expected, actual, t)
}

func Test2ShouldSay2(t *testing.T) {
	expected := "2"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(2)
	assert(expected, actual, t)
}

func Test3ShouldSayFizz(t *testing.T) {
	expected := "Fizz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(3)
	assert(expected, actual, t)
}

func Test4ShouldSay4(t *testing.T) {
	expected := "4"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(4)
	assert(expected, actual, t)
}

func Test5ShouldSayBuzz(t *testing.T) {
	expected := "Buzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(5)
	assert(expected, actual, t)
}

func Test6ShouldSayFizz(t *testing.T) {
	expected := "Fizz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(6)
	assert(expected, actual, t)
}

func Test7ShouldSay7(t *testing.T) {
	expected := "7"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(7)
	assert(expected, actual, t)
}

func Test8ShouldSay8(t *testing.T) {
	expected := "8"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(8)
	assert(expected, actual, t)
}

func Test9ShouldSayFizz(t *testing.T) {
	expected := "Fizz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(9)
	assert(expected, actual, t)
}

func Test10ShouldSayBuzz(t *testing.T) {
	expected := "Buzz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(10)
	assert(expected, actual, t)
}

func Test12ShouldSayFizz(t *testing.T) {
	expected := "Fizz"
	fizzBuzz := FizzBuzz{}
	actual := fizzBuzz.say(12)
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
