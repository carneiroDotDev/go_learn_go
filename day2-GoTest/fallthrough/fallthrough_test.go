package main

import (
	"testing"
)

func TestFallthrough(t *testing.T) {
	expected := "This is true! \nThis is false, but because of fallthrough, this thing executes! \n"
	actual := "This is true!"
	actualString := string(actual)
	if actualString != expected {
		t.Errorf("Expected %q, but got %q", expected, actualString)
	}
}