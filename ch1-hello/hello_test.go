package main

import "testing"

func TestHello(t *testing.T) {
	expected := "Hello, world"
	actual := hello()

	if expected != actual {
		t.Errorf("expected %q but was %q", expected, actual)
	}
}
