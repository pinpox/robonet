package robonet

import "testing"

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

func TestSigmoidFast(t *testing.T) {
	type args struct {
		x float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		//TODO add more test cases
		{"Test for zero", args{0.0}, 0.0},
		{"Test for zero", args{4}, 0.8}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SigmoidFast(tt.args.x); got != tt.want {
				t.Errorf("SigmoidFast() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareImages(t *testing.T) {
	type args struct {
		path1     string
		path2     string
		threshold float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareImages(tt.args.path1, tt.args.path2, tt.args.threshold); got != tt.want {
				t.Errorf("CompareImages() = %v, want %v", got, tt.want)
			}
		})
	}
}
