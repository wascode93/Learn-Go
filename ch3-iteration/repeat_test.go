package iteration

import "testing"

func TestRepeat(t *testing.T) {
	expected := "aaaaaa"
	actual := Repeat("a", 6)

	if expected != actual {
		t.Errorf("expected %q but was %q", expected, actual)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
