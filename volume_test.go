package robonet

import (
	"reflect"
	"testing"

	"github.com/gonum/matrix/mat64"
)

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

var testVol = NewWithData(3, 3, 3, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})
var testVolSquared = NewWithData(3, 3, 3, []float64{0, 1, 4, 9, 16, 25, 36, 49, 64, 81, 100, 121, 144, 169, 196, 225, 256, 289, 324, 361, 400, 441, 484, 529, 576, 625, 676})

// Check correctly sized volume is created
func TestNew(t *testing.T) {
	for _, v := range volumeSizes {
		vol := New(v[0], v[1], v[2])
		if !EqualNDim(vol.Shape(), v) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", vol.Shape())
		}
	}
}

// Check correctly sized volume is created
func TestNewRand(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewRand(v[0], v[1], v[2])
		if !EqualNDim(vol.Shape(), v) {
			t.Error("Expected ", v, ", got ", vol.Shape())
		}
	}
}

func TestVolume_Reflect(t *testing.T) {

	//Input
	testVol1 := NewWithData(3, 3, 3, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})
	testVol2 := NewWithData(3, 5, 2, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29})

	//Result
	testVol1Reflected := NewWithData(3, 3, 3, []float64{2, 1, 0, 5, 4, 3, 8, 7, 6, 11, 10, 9, 14, 13, 12, 17, 16, 15, 20, 19, 18, 23, 22, 21, 26, 25, 24})
	testVol2Reflected := NewWithData(3, 5, 2, []float64{4, 3, 2, 1, 0, 9, 8, 7, 6, 5, 14, 13, 12, 11, 10, 19, 18, 17, 16, 15, 24, 23, 22, 21, 20, 29, 28, 27, 26, 25})

	tests := []struct {
		name string
		vol1 Volume
		want Volume
	}{
		{"3 layers", testVol1, testVol1Reflected},
		{"2 layers", testVol2, testVol2Reflected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			vol.Reflect()
			if !vol.Equals(tt.want) {
				t.Errorf("Volume.Reflect() got = %v, want %v", vol, tt.want)
			}
		})
	}
}

func TestVolume_PointReflect(t *testing.T) {

	//Input
	testVol1 := NewWithData(3, 3, 3, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})
	testVol2 := NewWithData(3, 5, 2, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29})

	//Result
	testVol1PointReflected := NewWithData(3, 3, 3, []float64{0, 3, 6, 1, 4, 7, 2, 5, 8, 9, 12, 15, 10, 13, 16, 11, 14, 17, 18, 21, 24, 19, 22, 25, 20, 23, 26})

	testVol2PointReflected := NewWithData(5, 3, 2, []float64{0, 5, 10, 1, 6, 11, 2, 7, 12, 3, 8, 13, 4, 9, 14, 15, 20, 25, 16, 21, 26, 17, 22, 27, 18, 23, 28, 19, 24, 29})

	tests := []struct {
		name string
		vol1 Volume
		want Volume
	}{
		{"3 layers", testVol1, testVol1PointReflected},
		{"2 layers", testVol2, testVol2PointReflected},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			vol.PointReflect()
			if !vol.Equals(tt.want) {
				t.Errorf("Volume.PointReflect() got = %v, want %v", vol, tt.want)
			}
		})
	}

}

func TestVolume_SetAll(t *testing.T) {
	tests := []struct {
		name string
		in   Volume
		want Volume
		args Volume
	}{
		{"Set all to zero", NewRand(2, 2, 2), New(2, 2, 2), New(2, 2, 2)}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.in.SetAll(tt.args)
			if !tt.in.Equals(tt.want) {
				t.Errorf("Volume.SetAll() got = %v, want %v", tt.in, tt.want)
			}
		})
	}
}

// Check if Correct Dimensions are displayed
func TestVolume_Dims(t *testing.T) {
	tests := []struct {
		name  string
		vol   Volume
		want  int
		want1 int
		want2 int
	}{
		{"Cube", testVol, 3, 3, 3},
		{"All different", NewRand(1, 2, 3), 1, 2, 3},
		{"1- Sized", NewRand(1, 1, 1), 1, 1, 1},
		{"0- Sized", NewRand(0, 0, 0), 0, 0, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vs := tt.vol.Shape()
			if vs[0] != tt.want {
				t.Errorf("Volume.Dims() got = %v, want %v", vs[0], tt.want)
			}
			if vs[1] != tt.want1 {
				t.Errorf("Volume.Dims() got1 = %v, want %v", vs[1], tt.want1)
			}
			if vs[2] != tt.want2 {
				t.Errorf("Volume.Dims() got2 = %v, want %v", vs[2], tt.want2)
			}
		})
	}
}

func TestVolume_Apply(t *testing.T) {
	type args struct {
		kern    Kernel
		strideR int
		strideC int
	}
	tests := []struct {
		name string
		args args
		vol  Volume
		want Volume
	}{

	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			if vol.Apply(tt.args.kern, tt.args.strideR, tt.args.strideC); !vol.Equals(tt.want) {
				t.Errorf("Volume.Apply() got = %v, want %v", vol, tt.want)
			}
		})
	}
}

func TestNewFull(t *testing.T) {
	type args struct {
		r   int
		c   int
		d   int
		fil float64
	}
	tests := []struct {
		name string
		args args
		want Volume
	}{
		{"0-Filled", args{3, 3, 3, 0}, New(3, 3, 3)}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewFull(tt.args.r, tt.args.c, tt.args.d, tt.args.fil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_SubVolumePadded(t *testing.T) {

	subVol1 := NewWithData(3, 3, 3, []float64{0, 0, 0, 0, 0, 1, 0, 3, 4, 0, 0, 0, 0, 9, 10, 0, 12, 13, 0, 0, 0, 0, 18, 19, 0, 21, 22})
	subVol2 := NewWithData(3, 3, 3, []float64{0, 0, 0, 0, 1, 2, 3, 4, 5, 0, 0, 0, 9, 10, 11, 12, 13, 14, 0, 0, 0, 18, 19, 20, 21, 22, 23})
	subVol3 := NewWithData(3, 3, 3, []float64{0, 0, 0, 1, 2, 0, 4, 5, 0, 0, 0, 0, 10, 11, 0, 13, 14, 0, 0, 0, 0, 19, 20, 0, 22, 23, 0})

	subVol4 := NewWithData(3, 3, 3, []float64{0, 0, 1, 0, 3, 4, 0, 6, 7, 0, 9, 10, 0, 12, 13, 0, 15, 16, 0, 18, 19, 0, 21, 22, 0, 24, 25})

	//sub1 := data1
	//sub2 := data2
	//sub3 := data3

	//var subVol5 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	subVol6 := NewWithData(3, 3, 3, []float64{4, 5, 0, 7, 8, 0, 0, 0, 0, 13, 14, 0, 16, 17, 0, 0, 0, 0, 22, 23, 0, 25, 26, 0, 0, 0, 0})

	var subVol7 = NewWithData(1, 1, 3, []float64{0, 9, 18})

	sub1 := []float64{
		0, 0, 0, 0, 0,
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0,
		0, 0, 0, 0, 0,

		0, 0, 0, 0, 0,
		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0,
		0, 0, 0, 0, 0,

		0, 0, 0, 0, 0,
		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0,
		0, 0, 0, 0, 0}
	var subVol8 = NewWithData(5, 5, 3, sub1)

	sub1 = []float64{
		0, 0, 0,
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		0, 0, 0,

		0, 0, 0,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
		0, 0, 0,

		0, 0, 0,
		18, 19, 20,
		21, 22, 23,
		24, 25, 26,
		0, 0, 0}

	var subVol9 = NewWithData(5, 3, 3, sub1)

	sub1 = []float64{
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0,

		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0,

		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0}

	var subVol10 = NewWithData(3, 5, 3, sub1)

	sub1 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 2, 0, 0,
		0, 0, 3, 4, 5, 0, 0,
		0, 0, 6, 7, 8, 0, 0,
		0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0,
		0, 0, 9, 10, 11, 0, 0,
		0, 0, 12, 13, 14, 0, 0,
		0, 0, 15, 16, 17, 0, 0,
		0, 0, 0, 0, 0, 0, 0,

		0, 0, 0, 0, 0, 0, 0,
		0, 0, 18, 19, 20, 0, 0,
		0, 0, 21, 22, 23, 0, 0,
		0, 0, 24, 25, 26, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol11 = NewWithData(5, 7, 3, sub1)

	sub1 = []float64{
		0, 0, 1, 2, 0, 0, 0,
		0, 3, 4, 5, 0, 0, 0,
		0, 6, 7, 8, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,

		0, 9, 10, 11, 0, 0, 0,
		0, 12, 13, 14, 0, 0, 0,
		0, 15, 16, 17, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,

		0, 18, 19, 20, 0, 0, 0,
		0, 21, 22, 23, 0, 0, 0,
		0, 24, 25, 26, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol12 = NewWithData(5, 7, 3, sub1)

	type args struct {
		cR int
		cC int
		r  int
		c  int
	}
	tests := []struct {
		name string
		vol  Volume
		args args
		want Volume
	}{
		{"Default", testVol, args{0, 0, 3, 3}, subVol1},
		{"Default", testVol, args{0, 1, 3, 3}, subVol2},
		{"Default", testVol, args{0, 2, 3, 3}, subVol3},
		{"Default", testVol, args{1, 0, 3, 3}, subVol4},
		//{"Default", testVol, args{1, 1, 3, 3}, subVol5},
		{"Default", testVol, args{2, 2, 3, 3}, subVol6},
		{"Default", testVol, args{0, 0, 1, 1}, subVol7},
		{"Default", testVol, args{1, 1, 5, 5}, subVol8},
		{"Default", testVol, args{1, 1, 5, 3}, subVol9},
		{"Default", testVol, args{1, 1, 3, 5}, subVol10},
		{"Default", testVol, args{1, 1, 5, 7}, subVol11},
		{"Default", testVol, args{2, 2, 5, 7}, subVol12}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := testVol
			if got := vol.SubVolumePadded(tt.args.cR, tt.args.cC, tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Volume.SubVolumePadded() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Equals(t *testing.T) {
	tests := []struct {
		name string
		vol1 Volume
		vol2 Volume
		want bool
	}{
		{"Same", New(3, 3, 3), New(3, 3, 3), true},
		{"Different Dims", New(3, 3, 3), New(5, 5, 5), false},
		{"Different values", New(3, 3, 3), NewRand(3, 3, 3), false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if got := vol.Equals(tt.vol2); got != tt.want {
				t.Errorf("Volume.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_GetAt(t *testing.T) {
	type args struct {
		r int
		c int
		d int
	}
	tests := []struct {
		name string
		vol  Volume
		args args
		want float64
	}{
		{"First Element", testVol, args{0, 0, 0}, 0},
		{"First Element", testVol, args{2, 2, 2}, 26}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			if got := vol.GetAt(tt.args.r, tt.args.c, tt.args.d); got != tt.want {
				t.Errorf("Volume.GetAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_SetAt(t *testing.T) {

	res1 := NewWithData(3, 3, 3, []float64{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0})
	res2 := NewWithData(3, 3, 3, []float64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1})

	type args struct {
		r   int
		c   int
		d   int
		val float64
	}
	tests := []struct {
		name string
		vol1 Volume
		vol2 Volume
		args args
	}{
		{"First Element", New(3, 3, 3), res1, args{0, 0, 0, 1}},
		{"Last Element", New(3, 3, 3), res2, args{2, 2, 2, 1}}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1

			if vol.SetAt(tt.args.r, tt.args.c, tt.args.d, tt.args.val); !vol.Equals(tt.vol2) {
				t.Errorf("Volume.SetAt() = %v, want %v", vol, tt.vol2)
			}
		})
	}
}

func TestVolume_Print(t *testing.T) {
	type fields struct {
		Fields []mat64.Dense
	}
	tests := []struct {
		name string
		vol  Volume
	}{
		{"Default", testVol}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			vol.Print()
		})
	}
}

func TestVolume_Rows(t *testing.T) {
	tests := []struct {
		name string
		vol  Volume
		want int
	}{
		{"Normal", New(3, 3, 3), 3},
		{"1 Sized", New(1, 3, 3), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			if got := vol.Rows(); got != tt.want {
				t.Errorf("Volume.Rows() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Collumns(t *testing.T) {
	tests := []struct {
		name string
		vol  Volume
		want int
	}{
		{"Normal", New(3, 3, 3), 3},
		{"1 Sized", New(3, 1, 3), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			if got := vol.Collumns(); got != tt.want {
				t.Errorf("Volume.Collumns() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Depth(t *testing.T) {
	tests := []struct {
		name string
		vol  Volume
		want int
	}{
		{"Normal", New(3, 3, 3), 3},
		{"1 Sized", New(3, 3, 1), 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol
			if got := vol.Depth(); got != tt.want {
				t.Errorf("Volume.Depth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_EqualSize(t *testing.T) {
	type fields struct {
		Fields []mat64.Dense
	}
	type args struct {
		a Volume
	}
	tests := []struct {
		name string
		vol1 Volume
		vol2 Volume
		want bool
	}{
		{"Both equal", NewRand(3, 3, 3), NewRand(3, 3, 3), true},
		{"Different X", NewRand(1, 3, 3), NewRand(3, 3, 3), false},
		{"Different Y", NewRand(3, 1, 3), NewRand(3, 3, 3), false},
		{"Different Z", NewRand(3, 3, 1), NewRand(3, 3, 3), false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if got := vol.EqualSize(tt.vol2); got != tt.want {
				t.Errorf("Volume.EqualSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_MulElem(t *testing.T) {
	tests := []struct {
		name string
		vol1 Volume
		vol2 Volume
		want Volume
	}{
		{"All Zeros", New(3, 3, 3), New(3, 3, 3), New(3, 3, 3)},
		{"Zeros with random", New(3, 3, 3), NewRand(3, 3, 3), New(3, 3, 3)},
		{"Random with zeros", NewRand(3, 3, 3), New(3, 3, 3), New(3, 3, 3)},
		{"testVol squared", testVol, testVol, testVolSquared}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if vol.MulElem(tt.vol2); !vol.Equals(tt.want) {
				t.Errorf("Volume.MulElem = \n\n%v, want \n\n%v", vol, tt.want)
			}
		})
	}
}

func TestVolume_Max(t *testing.T) {

	test := NewWithData(3, 3, 3, []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26})

	tests := []struct {
		name string
		vol  Volume
		want float64
	}{
		{"Numbered", test, 26},
		{"All Zero", New(5, 5, 5), 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vol.Max(); got != tt.want {
				t.Errorf("Volume.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_Norm(t *testing.T) {

	tVol1 := testVol
	res1 := testVol
	res1.MulElem(NewFull(3, 3, 3, 100))
	res1.MulElem(NewFull(3, 3, 3, 1.0/26.0))

	tVol3 := New(3, 3, 2)
	tVol3.SetAt(0, 0, 0, 0)
	tVol3.SetAt(0, 1, 0, 1)
	tVol3.SetAt(0, 2, 0, 2)
	tVol3.SetAt(1, 0, 0, 3)
	tVol3.SetAt(1, 1, 0, 4)
	tVol3.SetAt(1, 2, 0, 5)
	tVol3.SetAt(2, 0, 0, 6)
	tVol3.SetAt(2, 1, 0, 7)
	tVol3.SetAt(2, 2, 0, 8)
	tVol3.SetAt(0, 0, 1, -0)
	tVol3.SetAt(0, 1, 1, -1)
	tVol3.SetAt(0, 2, 1, -2)
	tVol3.SetAt(1, 0, 1, -3)
	tVol3.SetAt(1, 1, 1, -4)
	tVol3.SetAt(1, 2, 1, -5)
	tVol3.SetAt(2, 0, 1, -6)
	tVol3.SetAt(2, 1, 1, -7)
	tVol3.SetAt(2, 2, 1, -8)

	res3 := New(3, 3, 2)

	res3.SetAt(0, 0, 0, 16)
	res3.SetAt(0, 1, 0, 18)
	res3.SetAt(0, 2, 0, 20)
	res3.SetAt(1, 0, 0, 22)
	res3.SetAt(1, 1, 0, 24)
	res3.SetAt(1, 2, 0, 26)
	res3.SetAt(2, 0, 0, 28)
	res3.SetAt(2, 1, 0, 30)
	res3.SetAt(2, 2, 0, 32)
	res3.SetAt(0, 0, 1, 16)
	res3.SetAt(0, 1, 1, 14)
	res3.SetAt(0, 2, 1, 12)
	res3.SetAt(1, 0, 1, 10)
	res3.SetAt(1, 1, 1, 8)
	res3.SetAt(1, 2, 1, 6)
	res3.SetAt(2, 0, 1, 4)
	res3.SetAt(2, 1, 1, 2)
	res3.SetAt(2, 2, 1, 0)

	tests := []struct {
		name string
		vol1 Volume
		max  float64
		want Volume
	}{
		{"positive only", tVol1, 100, res1},
		{"negative and postive", tVol3, 32, res3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if vol.Norm(tt.max); !vol.SimilarTo(tt.want, 0.01) {
				t.Errorf("Volume.SubVolume() = \n%v, want \n%v", vol, tt.want)
			}
		})
	}
}
