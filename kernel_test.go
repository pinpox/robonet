package robonet

import (
	"reflect"
	"testing"

	"github.com/gonum/matrix/mat64"
)

func TestKernelPointReflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1 := NewKernel(3, 3, 3)
	kern1.SetAll(testVol1)

	//Result
	layer1 = []float64{0, 3, 6, 1, 4, 7, 2, 5, 8}
	layer2 = []float64{9, 12, 15, 10, 13, 16, 11, 14, 17}
	layer3 = []float64{18, 21, 24, 19, 22, 25, 20, 23, 26}
	testVol1Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1Reflected := NewKernel(3, 3, 3)
	kern1Reflected.SetAll(testVol1Reflected)

	//Compare
	kern1.PointReflect()
	if !(kern1Reflected.Equals(kern1)) {

		t.Error("Expected", kern1Reflected, " got", kern1)
		kern1Reflected.Print()
		kern1.Print()
	}
}

func TestKernelReflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1 := NewKernel(3, 3, 3)
	kern1.SetAll(testVol1)

	//Result
	layer1 = []float64{2, 1, 0, 5, 4, 3, 8, 7, 6}
	layer2 = []float64{11, 10, 9, 14, 13, 12, 17, 16, 15}
	layer3 = []float64{20, 19, 18, 23, 22, 21, 26, 25, 24}
	testVol1Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1Reflected := NewKernel(3, 3, 3)
	kern1Reflected.SetAll(testVol1Reflected)

	//Compare
	kern1.Reflect()
	if !(kern1Reflected.Equals(kern1)) {

		t.Error("Reflect () Expected", kern1Reflected, " got", kern1)
		kern1Reflected.Print()
		kern1.Print()
	}
}

func TestNewKernel(t *testing.T) {
	type args struct {
		r int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want Kernel
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKernel(tt.args.r, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKernel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKernel_Equals(t *testing.T) {
	type fields struct {
		Volume Volume
	}
	type args struct {
		in Kernel
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kern := &Kernel{
				Volume: tt.fields.Volume,
			}
			if got := kern.Equals(tt.args.in); got != tt.want {
				t.Errorf("Kernel.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKernelRandom(t *testing.T) {
	type args struct {
		r int
		c int
		d int
	}
	tests := []struct {
		name string
		args args
		want Kernel
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKernelRandom(tt.args.r, tt.args.c, tt.args.d); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKernelRandom() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewKernelFilled(t *testing.T) {
	type args struct {
		r   int
		c   int
		d   int
		fil float64
	}
	tests := []struct {
		name string
		args args
		want Kernel
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewKernelFilled(tt.args.r, tt.args.c, tt.args.d, tt.args.fil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewKernelFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKernel_Apply(t *testing.T) {

	layer1 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	layer2 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	layer3 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker1 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker2 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker3 := []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}
	kernVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, ker1), *mat64.NewDense(3, 3, ker2), *mat64.NewDense(3, 3, ker3)}}
	kern1 := NewKernel(3, 3, 3)
	kern1.SetAll(kernVol1)

	layer1 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	layer2 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	layer3 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker1 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker2 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker3 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	testVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}
	kernVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, ker1), *mat64.NewDense(3, 3, ker2), *mat64.NewDense(3, 3, ker3)}}
	kern2 := NewKernel(3, 3, 3)
	kern2.SetAll(kernVol2)

	layer1 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	layer2 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	layer3 = []float64{1, 1, 1, 1, 1, 1, 1, 1, 1}
	ker1 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker2 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker3 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	testVol3 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}
	kernVol3 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, ker1), *mat64.NewDense(3, 3, ker2), *mat64.NewDense(3, 3, ker3)}}
	kern3 := NewKernel(3, 3, 3)
	kern3.SetAll(kernVol3)

	layer1 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	layer2 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	layer3 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker1 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker2 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	ker3 = []float64{0, 0, 0, 0, 0, 0, 0, 0, 0}
	testVol4 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}
	kernVol4 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, ker1), *mat64.NewDense(3, 3, ker2), *mat64.NewDense(3, 3, ker3)}}
	kern4 := NewKernel(3, 3, 3)
	kern4.SetAll(kernVol4)

	tests := []struct {
		name string
		kern Kernel
		vol  Volume
		want float64
	}{
		{"vol ones kern ones", kern1, testVol1, 1},
		{"vol zeros kern ones", kern2, testVol2, 0},
		{"vol ones kern zeros", kern3, testVol3, 0},
		{"vol zeros kern zeros", kern4, testVol4, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kern := tt.kern
			if got := kern.Apply(tt.vol); Round(got, 10) != tt.want {
				t.Errorf("Kernel.Apply() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKernel_Elems(t *testing.T) {
	tests := []struct {
		name string
		kern Kernel
		want int
	}{
		{"3x3x3", NewKernelRandom(3, 3, 3), 27},
		{"3x3x1", NewKernelRandom(3, 3, 1), 9},
		{"3x5x1", NewKernelRandom(3, 5, 1), 15},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			kern := tt.kern
			if got := kern.Elems(); got != tt.want {
				t.Errorf("Kernel.Elems() = %v, want %v", got, tt.want)
			}
		})
	}
}
