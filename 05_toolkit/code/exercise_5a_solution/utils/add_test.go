package utils

import "testing"

func TestAdd(t *testing.T) {
	expected := 4
	actual := Add(2, 2)
	
	if actual != expected {
		t.Errorf("Sum was incorrect! Expected: %d, Actual: %d", expected, actual)
	}
}
