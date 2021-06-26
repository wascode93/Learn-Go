package integers

import "testing"

func TestAdder(t *testing.T) {
	expected := 4
	actual := Add(2, 2)

	if expected != actual {
		t.Errorf("expected %d but was %d", expected, actual)
	}
}
