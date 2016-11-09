package robonet

import "testing"
import "github.com/gonum/matrix/mat64"

var volumeSizes = [][]int{
	{1, 1, 1},
	{3, 3, 3},
	{5, 5, 5},
	{5, 5, 1},
	{5, 1, 5},
	{1, 1, 5},
	{5, 1, 5},
	{1, 2, 3},
	{3, 2, 1},
}

func TestDims(t *testing.T) {
	for _, v := range volumeSizes {
		vol := new(rNVolume)

		for i := 0; i < v[2]; i++ {
			vol.Fields = append(vol.Fields, *mat64.NewDense(v[0], v[1], nil))
		}

		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}

}

// Check correctly sized volume is created
func TestNewRNVolume(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewRNVolume(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}

// Check correctly sized volume is created
func TestNewRNVolumeRandom(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewRNVolumeRandom(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}
