package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, Waseem"
	actual := hello("Waseem")

	if expected != actual {
		t.Errorf("expected %q but was %q", expected, actual)
	}
}
