package fizzbuzz

import (
	"testing"
)

func Test1ShouldSay1(t *testing.T) {
	expected := "1"

	actual := Say(1)

	if expected != actual {
		t.Fatal("Expected say", expected, "but actually said", actual)
	}
}
