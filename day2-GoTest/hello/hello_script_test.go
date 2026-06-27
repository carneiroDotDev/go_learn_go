package main

import "testing"

func TestHello(t *testing.T) {

	checkCorrectMessage := func (t *testing.T, actual string, expected string)  {
		t.Helper()
		if actual != expected {
			t.Errorf("Expected %q, but got %q", expected, actual)
		}
	}

	t.Run("Test with name", func(t *testing.T) {
		actual := hello("Luiz", "")
		expected := "Hello, Luiz!\n"
		checkCorrectMessage(t, actual, expected)
	})

	t.Run("Test without name", func(t *testing.T) {
		actual := hello("", "")
		expected := "Hello, World!\n"
		checkCorrectMessage(t, actual, expected)
	})

	t.Run("in Spanish", func(t *testing.T) {
		actual := hello("Elodie", "es")
        expected := "Hola, Elodie!\n"
        checkCorrectMessage(t, actual, expected)
    })

	t.Run("in French", func(t *testing.T) {
		actual := hello("Gael", "fr")
        expected := "Bonjour, Gael!\n"
        checkCorrectMessage(t, actual, expected)
    })
}