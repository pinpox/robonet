package robonet

import "testing"

func TestAddKernel(t *testing.T) {

	layer := ConvLayer{}

	layer.AddKernel(NewKernel(3, 3, 3), 1, 1)
	layer.AddKernel(NewKernel(3, 3, 3), 2, 1)
	layer.AddKernel(NewKernel(3, 3, 3), 1, 4)

	if len(layer.Kernels()) != 3 {
		t.Error("Expected ", 3, ", got ", len(layer.Kernels()))
	}
}

func TestNormCalculate(t *testing.T) {
	//TODO
}

func TestConvCalculate(t *testing.T) {
	//TODO
}

func TestPoolCalculate(t *testing.T) {
	//TODO
}
func TestFCCalculate(t *testing.T) {
	//TODO
}

func TestReluCalculate(t *testing.T) {
	//TODO
}
