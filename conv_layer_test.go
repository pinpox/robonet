package robonet

import (
	"reflect"
	"testing"
)

func TestConvLayer_AddKernel(t *testing.T) {

	lay1 := *new(ConvLayer)
	lay2 := *new(ConvLayer)

	ker := NewKernelRandom(3, 3, 3)

	lay1.AddKernel(ker, 2, 2)
	lay1.AddKernel(ker, 2, 2)
	lay1.AddKernel(ker, 2, 2)

	lay2.AddKernel(ker, 2, 2)
	lay2.AddKernel(ker, 2, 2)
	lay2.AddKernel(ker, 2, 2)
	lay2.AddKernel(ker, 2, 2)
	lay2.AddKernel(ker, 2, 2)
	lay2.AddKernel(ker, 2, 2)

	tests := []struct {
		name string
		want int
		lay  ConvLayer
	}{
		{"3 Kernels", 3, lay1},
		{"6 kernels", 6, lay2}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := tt.lay
			if got := len(l.Kernels()); got != tt.want {
				t.Errorf("ConvLayer.AddKernel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvLayer_Kernels(t *testing.T) {

	lay := *new(ConvLayer)
	lay.AddKernel(NewKernel(3, 1, 2), 1, 1)
	lay.AddKernel(NewKernel(1, 1, 1), 1, 1)
	lay.AddKernel(NewKernel(3, 3, 3), 1, 1)
	kers := []Kernel{NewKernel(3, 1, 2), NewKernel(1, 1, 1), NewKernel(3, 3, 3)}

	tests := []struct {
		name string
		lay  ConvLayer
		want []Kernel
	}{
		{"Default", lay, kers}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lay
			if got := l.Kernels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvLayer.Kernels() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestConvLayer_Calculate(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
		kernels     []Kernel
		strideR     int
		strideC     int
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ConvLayer{
				LayerFields: tt.fields.LayerFields,
				kernels:     tt.fields.kernels,
				strideR:     tt.fields.strideR,
				strideC:     tt.fields.strideC,
			}
			l.Calculate()
		})
	}
}
