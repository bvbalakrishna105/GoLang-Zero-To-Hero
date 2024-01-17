package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 3)

	if result != 5 {
		t.Errorf("Add(2,3) = %v want 5", result)
	}
}

func TestMul(t *testing.T) {
	result := Mul(2, 3)

	if result != 6 {
		t.Errorf("Mul(2,3) = %v want 6", result)
	}
}

func TestDiv(t *testing.T) {
	result := Div(4, 2)

	if result != 2 {
		t.Errorf("Div(4,2) = %v want 2", result)
	}
}

func TestSub(t *testing.T) {
	result := Sub(3, 2)

	if result != 1 {
		t.Errorf("Sub(3,2) = %v want 1", result)
	}
}
