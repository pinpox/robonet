package robonet

import (
	"reflect"
	"testing"
)

func TestConvLayer_AddKernel(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
		kernels     []Kernel
		stridesR    []int
		stridesC    []int
	}
	type args struct {
		kern    Kernel
		strideR int
		strideC int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &ConvLayer{
				LayerFields: tt.fields.LayerFields,
				kernels:     tt.fields.kernels,
				stridesR:    tt.fields.stridesR,
				stridesC:    tt.fields.stridesC,
			}
			l.AddKernel(tt.args.kern, tt.args.strideR, tt.args.strideC)
		})
	}
}

func TestConvLayer_Calculate(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
		kernels     []Kernel
		stridesR    []int
		stridesC    []int
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
				stridesR:    tt.fields.stridesR,
				stridesC:    tt.fields.stridesC,
			}
			l.Calculate()
		})
	}
}

func TestConvLayer_Kernels(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
		kernels     []Kernel
		stridesR    []int
		stridesC    []int
	}
	tests := []struct {
		name   string
		fields fields
		want   []Kernel
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := ConvLayer{
				LayerFields: tt.fields.LayerFields,
				kernels:     tt.fields.kernels,
				stridesR:    tt.fields.stridesR,
				stridesC:    tt.fields.stridesC,
			}
			if got := l.Kernels(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ConvLayer.Kernels() = %v, want %v", got, tt.want)
			}
		})
	}
}
