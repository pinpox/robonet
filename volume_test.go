package robonet

import "testing"

func TestNewRNVolume(t *testing.T) {
	v := NewRNVolume(3, 3, 3)
	i1, i2, i3 := v.Dims()
	if !Equal3Dim(i1, i2, i3, 3, 3, 3) {
		t.Error("Expected 3x3x3, got ", i1, i2, i3)
	}
}
