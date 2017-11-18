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
