package robonet

import (
	"image"
	"reflect"
	"testing"
)

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

func TestVolumeFromTIFF(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VolumeFromTIFF(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolumeFromTIFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolumeFromJPEG(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VolumeFromJPEG(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolumeFromJPEG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSaveVolumeToTIFF(t *testing.T) {
	type args struct {
		path string
		vol  Volume
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveVolumeToTIFF(tt.args.path, tt.args.vol)
		})
	}
}

func TestSaveVolumeToJPEG(t *testing.T) {
	type args struct {
		path string
		vol  Volume
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SaveVolumeToJPEG(tt.args.path, tt.args.vol)
		})
	}
}

func TestImageToVolume(t *testing.T) {
	type args struct {
		img image.Image
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ImageToVolume(tt.args.img); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ImageToVolume() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolumeToImage(t *testing.T) {
	type args struct {
		vol Volume
	}
	tests := []struct {
		name string
		args args
		want image.Image
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := VolumeToImage(tt.args.vol); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("VolumeToImage() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareJPEG(t *testing.T) {
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
		{"Same image thresh 0", args{"images/grey0.jpeg", "images/grey0.jpeg", 0}, true},
		{"Same image thresh 10", args{"images/grey0.jpeg", "images/grey0.jpeg", 10}, true},
		{"Different Images thresh 0", args{"images/grey1.jpeg", "images/grey0.jpeg", 0}, false},
		{"Different Images thresh 30", args{"images/grey30.jpeg", "images/grey0.jpeg", 30}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareJPEG(tt.args.path1, tt.args.path2, tt.args.threshold); got != tt.want {
				t.Errorf("CompareJPEG() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareTIFF(t *testing.T) {
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
		{"Same image thresh 0", args{"images/grey0.tiff", "images/grey0.tiff", 0}, true},
		{"Same image thresh 10", args{"images/grey0.tiff", "images/grey0.tiff", 10}, true},
		{"Different Images thresh 0", args{"images/grey1.tiff", "images/grey0.tiff", 0}, false},
		{"Different Images thresh 30", args{"images/grey30.tiff", "images/grey0.tiff", 30}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTIFF(tt.args.path1, tt.args.path2, tt.args.threshold); got != tt.want {
				t.Errorf("CompareTIFF() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRound(t *testing.T) {
	type args struct {
		val    float64
		places int
	}
	tests := []struct {
		name       string
		args       args
		wantNewVal float64
	}{
		{"Round up positive", args{0.006, 2}, 0.01},
		{"Round up positive", args{0.005, 2}, 0.01},
		{"Round down positive", args{0.004, 2}, 0.00},
		{"Round up negative", args{-0.006, 2}, -0.01},
		{"Round up negative", args{-0.005, 2}, -0.01},
		{"Round down negative", args{-0.004, 2}, -0.01},
		{"large places", args{0.004, 100}, 0.004},
		{"large places", args{0.0000000000000000000000000000000004, 34}, 0.0000000000000000000000000000000004},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotNewVal := Round(tt.args.val, tt.args.places); gotNewVal != tt.wantNewVal {
				t.Errorf("Round() = %v, want %v", gotNewVal, tt.wantNewVal)
			}
		})
	}
}
