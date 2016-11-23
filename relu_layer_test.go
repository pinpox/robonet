package robonet

import "testing"

func TestReluLayer_Calculate(t *testing.T) {
	tests := []struct {
		name  string
		lay   ReluLayer
		input Volume
		want  Volume
	}{
		{"Zero input", *new(ReluLayer), New(3, 3, 3), New(3, 3, 3)},
		{"4 filled input", *new(ReluLayer), NewFull(3, 3, 3, 4), NewFull(3, 3, 3, 0.8)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lay := tt.lay
			lay.input = tt.input
			if lay.Calculate(); !lay.output.Equals(tt.want) {
				t.Errorf("ReluLayer.Calculate() = %v, want %v", lay.output, tt.want)
			}
		})
	}
}
