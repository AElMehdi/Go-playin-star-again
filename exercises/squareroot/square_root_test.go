package squareroot

import (
	"strings"
	"testing"
)

func TestSquareRootOfValueOne(t *testing.T) {
	result, err := Sqrt(1)

	if result != 1 || err != nil {
		t.Errorf("Expecting %f to be %d", result, 1)
	}
}

func TestSquareRootValueFour(t *testing.T) {
	result, err := Sqrt(4)
	if result != 2 || err != nil {
		t.Errorf("Expecting %f to be %d", result, 2)
	}
}

func TestSquareRootNegativeValueError(t *testing.T) {
	_, err := Sqrt(-2)

	if err == nil {
		t.Errorf("Expecting an error on negative value %d", -2)

		negativeSqrt, ok := err.(ErrNegativeSqrt)
		if !ok {
			t.Errorf("Expecting an error on negative value of type NegativeSqrt error")

			errorMessage := "Cannot Sqrt negative number"
			if !strings.Contains(negativeSqrt.Error(), errorMessage) {
				t.Errorf("Expecting the error message to contain: %s", errorMessage)
			}
		}
	}
}
