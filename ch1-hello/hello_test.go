package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, expected, actual string) {
		t.Helper()
		if expected != actual {
			t.Errorf("expected %q but was %q", expected, actual)
		}
	}

	t.Run("Saying hello to people", func(t *testing.T) {
		expected := "Hello, Waseem"
		actual := hello("Waseem")
		assertCorrectMessage(t, expected, actual)
	})

	t.Run("Say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		expected := "Hello, world"
		actual := hello("")
		assertCorrectMessage(t, expected, actual)
	})
}
