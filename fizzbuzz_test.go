package fizzbuzz

import (
	"testing"
)

func Test1ShouldSay1(t *testing.T) {
	expected := "1"

	actual := say(1)

	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
	}
}

func Test2ShouldSay2(t *testing.T) {
	expected := "2"

	actual := say(2)

	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
	}
}
