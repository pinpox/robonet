package robonet

import "testing"

func TestNormLayer_Calculate(t *testing.T) {
	tests := []struct {
		name  string
		lay   NormLayer
		val   float64
		input Volume
	}{
		{"Zero input", *new(NormLayer), 0.0, New(3, 3, 3)},
		{"100 filled input to 50", *new(NormLayer), 50.0, NewFull(3, 3, 3, 100)},
		{"100 filled input to 1", *new(NormLayer), 1.0, NewFull(3, 3, 3, 100)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lay := tt.lay
			lay.input = tt.input
			if lay.Calculate(); lay.output.Max() > tt.val {
				t.Errorf("NormLayer.Calculate().Max() = %v, want %v", lay.output.Max(), tt.val)
			}
		})
	}
}
