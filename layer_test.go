package robonet

import "testing"

func TestAddKernel(t *testing.T) {

	layer := RNConvLayer{}

	layer.AddKernel(NewKernel(3, 3, 3))
	layer.AddKernel(NewKernel(3, 3, 3))
	layer.AddKernel(NewKernel(3, 3, 3))

	if len(layer.Kernels) != 3 {
		t.Error("Expected ", 3, ", got ", len(layer.Kernels))
	}
}

func TestCalculate(t *testing.T) {
	//TODO
}
