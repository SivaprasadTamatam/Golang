package maths

import "testing"

func TestAdd(t *testing.T) {
	res := Add(2, 3)
	if res != 5 {
		t.Error("Add Test Failed")
	}
}

func TestMul(t *testing.T) {
	res := Mul(2, 3)
	if res != 6 {
		t.Error("Mul Test Failed")
	}
}
