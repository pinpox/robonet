package robonet

import "testing"

func TestSigmoidFast(t *testing.T) {

	if SigmoidFast(4) != 0.8 {
		t.Error("Expected ", 0.8, ", got ", SigmoidFast(4))
	}
}

func TestEqual3Dim(t *testing.T) {
	if !Equal3Dim(1, 2, 3, 1, 2, 3) {
		t.Error("Expected true, got", Equal3Dim(1, 2, 3, 1, 2, 3))
	}
	if Equal3Dim(1, 1, 3, 1, 2, 3) {
		t.Error("Expected false, got", Equal3Dim(1, 1, 3, 1, 2, 3))
	}
}

func TestOdd3Dim(t *testing.T) {
	if !Odd3Dim(1, 1, 1) {
		t.Error("Expected false, got", Odd3Dim(1, 1, 1))
	}
	if Odd3Dim(2, 2, 1) {
		t.Error("Expected true, got", Odd3Dim(2, 2, 1))
	}
	if !Odd3Dim(2, 1, 1) {
		t.Error("Expected false, got", Odd3Dim(2, 1, 1))
	}
}
