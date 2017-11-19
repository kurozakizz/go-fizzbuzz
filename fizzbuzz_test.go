package fizzbuzz

import (
	"testing"
)

func assert(expected string, actual string, t *testing.T) {
	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
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
