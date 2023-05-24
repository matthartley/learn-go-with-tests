package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a", 20)
	expected := "aaaaaaaaaaaaaaaaaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}
