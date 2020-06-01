package main

import (
	"fmt"
	"testing"
)

// TestPi is a regression test for the Pi function. Run: go test calculatePi_test.go calculatePi.go
func TestPi(t *testing.T) {
	tests := map[uint]string{
		1:   "4",
		2:   "3",
		60:  "3.14159265358979324",
		600: "3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105557",
		//    3.14159265358979323846264338327950288419716939937510582097494459230781640628620899862803482534211706798214808651328230664709384460955058223172535940812848111745028410270193852110555964462294895493038196
		// 	  ^ First 200 digits of Pi. Source: https://www.piday.org/million/
	}
	for input, expected := range tests {
		t.Run("", func(t *testing.T) {
			precision = input
			actual := Pi(precision)
			if fmt.Sprintf("%v", actual) != expected {
				t.Errorf("Pi(%d) = %v; want %s", input, actual, expected)
			}
		})
	}
}
