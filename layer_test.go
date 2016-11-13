package robonet

import "testing"

// Check if Correct Dimensions are displayed
//func TestDims(t *testing.T) {
//if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
//t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
//}
//}

func TestAddKernel(t *testing.T) {

	layer := new(RNConvLayer)

	kernel := new(Kernel)

	layer.AddKernel(kernel)
	layer.AddKernel(kernel)
	layer.AddKernel(kernel)

	if len(layer.Kernels) != 3 {
		t.Error("Expected ", 3, ", got ", len(layer.Kernels))
	}
}

func TestCalculate(t *testing.T) {
	//TODO
}
