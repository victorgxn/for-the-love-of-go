package calculator

import (
	"errors"
	"fmt"
	"math"
)

func Add(a, b float64) float64 {
	return a + b
}

func Subtract(a, b float64) float64 {
	return a - b
}

func Multiply(a, b float64) float64 {
	return a * b
}

func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero not allowed")
	}
	return a / b, nil
}

func Sqrt(input float64) (float64, error) {
	if input < 0 {
		return 0, fmt.Errorf("square root of negative number not allowed: %f", input)
	}
	return math.Sqrt(input), nil
}
