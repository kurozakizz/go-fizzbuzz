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
	actual := say(1)
	assert(expected, actual, t)
}

func Test2ShouldSay2(t *testing.T) {
	expected := "2"
	actual := say(2)
	assert(expected, actual, t)
}

func Test3ShouldSayFizz(t *testing.T) {
	expected := "Fizz"
	actual := say(3)
	assert(expected, actual, t)
}

func Test4ShouldSay4(t *testing.T) {
	expected := "4"
	actual := say(4)
	assert(expected, actual, t)
}

func Test5ShouldSayBuzz(t *testing.T) {
	expected := "Buzz"
	actual := say(5)
	assert(expected, actual, t)
}
