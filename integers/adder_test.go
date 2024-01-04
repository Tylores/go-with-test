package integers

import "testing"

func TestAddr(t *testing.T) {
	sum := Add()
	expected := 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
