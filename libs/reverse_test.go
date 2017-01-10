package libs

import (
	"testing"
)

func TestReverse(t *testing.T) {

	expected := "olleH"
	actual := Reverse("Hello")

	if expected != actual {
		t.Error("Strings are not equal")
	}
}

func TestReverse2(t *testing.T) {
	expected := ""
	actual := Reverse("")

	if expected != actual {
		t.Error("Strings are not equal")
	}
}
