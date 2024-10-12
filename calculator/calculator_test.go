package calculator_test

import (
	"calculator"
	"testing"
)

func TestSubtract(t *testing.T) {
	//Arrange	
	t.Parallel()
	var want float64 = 2
	got := calculator.Subtract(2, 4)
	if want != got {
		t.Errorf("want %f, got %f", want, got)
	}
}
