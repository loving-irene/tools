package simplemath

import "testing"

func TestSqrt(t *testing.T) {
	v := Sqrt(4)
	if v != 2 {
		t.Errorf("Result %d,should be 2", v)
	}
}
