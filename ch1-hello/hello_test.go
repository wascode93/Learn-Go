package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %q but was %q", expected, actual)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		expected := "Hello, Waseem"
		actual := hello("Waseem", "")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		expected := "Hello, world"
		actual := hello("", "")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("in Spanish", func(t *testing.T) {
		expected := "Hola, Elodie"
		actual := hello("Elodie", "Spanish")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("in French", func(t *testing.T) {
		expected := "Bonjour, Elodie"
		actual := hello("Elodie", "French")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("in Arabic", func(t *testing.T) {
		expected := "أهلاً, وسيم"
		actual := hello("وسيم", "Arabic")
		assertCorrectMessage(t, expected, actual)
	})
}
