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

var data1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
var data2 = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
var data3 = []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
var testVol = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, data1), *mat64.NewDense(3, 3, data2), *mat64.NewDense(3, 3, data3)}}

var data4 = []float64{0, 1, 4, 9, 16, 25, 36, 49, 64}
var data5 = []float64{81, 100, 121, 144, 169, 196, 225, 256, 289}
var data6 = []float64{324, 361, 400, 441, 484, 529, 576, 625, 676}

var testVolSquared = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, data4), *mat64.NewDense(3, 3, data5), *mat64.NewDense(3, 3, data6)}}

// Check correctly sized volume is created
func TestNewVolume(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewVolume(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}

// Check correctly sized volume is created
func TestNewVolumeRandom(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewVolumeRandom(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}

func TestVolume_Reflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	layer2 = []float64{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	testVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}

	//Result
	layer1 = []float64{2, 1, 0, 5, 4, 3, 8, 7, 6}
	layer2 = []float64{11, 10, 9, 14, 13, 12, 17, 16, 15}
	layer3 = []float64{20, 19, 18, 23, 22, 21, 26, 25, 24}
	testVol1Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{4, 3, 2, 1, 0, 9, 8, 7, 6, 5, 14, 13, 12, 11, 10}
	layer2 = []float64{19, 18, 17, 16, 15, 24, 23, 22, 21, 20, 29, 28, 27, 26, 25}
	testVol2Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}

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
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	layer2 = []float64{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	testVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}

	//Result
	layer1 = []float64{0, 3, 6, 1, 4, 7, 2, 5, 8}
	layer2 = []float64{9, 12, 15, 10, 13, 16, 11, 14, 17}
	layer3 = []float64{18, 21, 24, 19, 22, 25, 20, 23, 26}
	testVol1PointReflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{0, 5, 10, 1, 6, 11, 2, 7, 12, 3, 8, 13, 4, 9, 14}
	layer2 = []float64{15, 20, 25, 16, 21, 26, 17, 22, 27, 18, 23, 28, 19, 24, 29}
	testVol2PointReflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 3, layer1), *mat64.NewDense(5, 3, layer2)}}

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
		{"Set all to zero", NewVolumeRandom(2, 2, 2), NewVolume(2, 2, 2), NewVolume(2, 2, 2)}}
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
		{"All different", NewVolumeRandom(1, 2, 3), 1, 2, 3},
		{"1- Sized", NewVolumeRandom(1, 1, 1), 1, 1, 1},
		{"0- Sized", NewVolumeRandom(0, 0, 0), 0, 0, 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2 := tt.vol.Dims()
			if got != tt.want {
				t.Errorf("Volume.Dims() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Volume.Dims() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("Volume.Dims() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}

func TestVolume_Apply(t *testing.T) {
	type fields struct {
		Fields []mat64.Dense
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
			vol := &Volume{
				Fields: tt.fields.Fields,
			}
			vol.Apply(tt.args.kern, tt.args.strideR, tt.args.strideC)
		})
	}
}

func TestNewVolumeFilled(t *testing.T) {
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
		{"0-Filled", args{3, 3, 3, 0}, NewVolume(3, 3, 3)}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewVolumeFilled(tt.args.r, tt.args.c, tt.args.d, tt.args.fil); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVolumeFilled() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_SubVolumePadded(t *testing.T) {

	var sub1 = []float64{
		0, 0, 0,
		0, 0, 1,
		0, 3, 4}

	var sub2 = []float64{
		0, 0, 0,
		0, 9, 10,
		0, 12, 13}

	var sub3 = []float64{
		0, 0, 0,
		0, 18, 19,
		0, 21, 22}

	var subVol1 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}
	sub1 = []float64{
		0, 0, 0,
		0, 1, 2,
		3, 4, 5}

	sub2 = []float64{
		0, 0, 0,
		9, 10, 11,
		12, 13, 14}

	sub3 = []float64{
		0, 0, 0,
		18, 19, 20,
		21, 22, 23}

	var subVol2 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		0, 0, 0,
		1, 2, 0,
		4, 5, 0}

	sub2 = []float64{
		0, 0, 0,
		10, 11, 0,
		13, 14, 0}

	sub3 = []float64{
		0, 0, 0,
		19, 20, 0,
		22, 23, 0}

	var subVol3 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		0, 0, 1,
		0, 3, 4,
		0, 6, 7}

	sub2 = []float64{
		0, 9, 10,
		0, 12, 13,
		0, 15, 16}

	sub3 = []float64{
		0, 18, 19,
		0, 21, 22,
		0, 24, 25}

	var subVol4 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = data1
	sub2 = data2
	sub3 = data3

	var subVol5 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		4, 5, 0,
		7, 8, 0,
		0, 0, 0}

	sub2 = []float64{
		13, 14, 0,
		16, 17, 0,
		0, 0, 0}

	sub3 = []float64{
		22, 23, 0,
		25, 26, 0,
		0, 0, 0}

	var subVol6 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{0}
	sub2 = []float64{9}
	sub3 = []float64{18}
	var subVol7 = Volume{Fields: []mat64.Dense{*mat64.NewDense(1, 1, sub1), *mat64.NewDense(1, 1, sub2), *mat64.NewDense(1, 1, sub3)}}

	sub1 = []float64{
		0, 0, 0, 0, 0,
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0,
		0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 0, 0, 0, 0,
		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0,
		0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 0, 0, 0, 0,
		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0,
		0, 0, 0, 0, 0}
	var subVol8 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 5, sub1), *mat64.NewDense(5, 5, sub2), *mat64.NewDense(5, 5, sub3)}}

	sub1 = []float64{
		0, 0, 0,
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		0, 0, 0}

	sub2 = []float64{
		0, 0, 0,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
		0, 0, 0}

	sub3 = []float64{
		0, 0, 0,
		18, 19, 20,
		21, 22, 23,
		24, 25, 26,
		0, 0, 0}

	var subVol9 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 3, sub1), *mat64.NewDense(5, 3, sub2), *mat64.NewDense(5, 3, sub3)}}

	sub1 = []float64{
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0}

	sub2 = []float64{
		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0}

	sub3 = []float64{
		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0}

	var subVol10 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, sub1), *mat64.NewDense(3, 5, sub2), *mat64.NewDense(3, 5, sub3)}}

	sub1 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 2, 0, 0,
		0, 0, 3, 4, 5, 0, 0,
		0, 0, 6, 7, 8, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 9, 10, 11, 0, 0,
		0, 0, 12, 13, 14, 0, 0,
		0, 0, 15, 16, 17, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 18, 19, 20, 0, 0,
		0, 0, 21, 22, 23, 0, 0,
		0, 0, 24, 25, 26, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol11 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 7, sub1), *mat64.NewDense(5, 7, sub2), *mat64.NewDense(5, 7, sub3)}}

	sub1 = []float64{
		0, 0, 1, 2, 0, 0, 0,
		0, 3, 4, 5, 0, 0, 0,
		0, 6, 7, 8, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 9, 10, 11, 0, 0, 0,
		0, 12, 13, 14, 0, 0, 0,
		0, 15, 16, 17, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 18, 19, 20, 0, 0, 0,
		0, 21, 22, 23, 0, 0, 0,
		0, 24, 25, 26, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol12 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 7, sub1), *mat64.NewDense(5, 7, sub2), *mat64.NewDense(5, 7, sub3)}}

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
		{"Default", testVol, args{1, 1, 3, 3}, subVol5},
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
		{"Same", NewVolume(3, 3, 3), NewVolume(3, 3, 3), true},
		{"Different Dims", NewVolume(3, 3, 3), NewVolume(5, 5, 5), false},
		{"Different values", NewVolume(3, 3, 3), NewVolumeRandom(3, 3, 3), false}}
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

	res1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, []float64{1, 0, 0, 0, 0, 0, 0, 0, 0}), *mat64.NewDense(3, 3, nil), *mat64.NewDense(3, 3, nil)}}
	res2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, nil), *mat64.NewDense(3, 3, nil), *mat64.NewDense(3, 3, []float64{0, 0, 0, 0, 0, 0, 0, 0, 1})}}

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
		{"First Element", NewVolume(3, 3, 3), res1, args{0, 0, 0, 1}},
		{"Last Element", NewVolume(3, 3, 3), res2, args{2, 2, 2, 1}}}
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
		{"Normal", NewVolume(3, 3, 3), 3},
		{"1 Sized", NewVolume(1, 3, 3), 1},
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
		{"Normal", NewVolume(3, 3, 3), 3},
		{"1 Sized", NewVolume(3, 1, 3), 1},
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
		{"Normal", NewVolume(3, 3, 3), 3},
		{"1 Sized", NewVolume(3, 3, 1), 1},
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
		{"Both equal", NewVolumeRandom(3, 3, 3), NewVolumeRandom(3, 3, 3), true},
		{"Different X", NewVolumeRandom(1, 3, 3), NewVolumeRandom(3, 3, 3), false},
		{"Different Y", NewVolumeRandom(3, 1, 3), NewVolumeRandom(3, 3, 3), false},
		{"Different Z", NewVolumeRandom(3, 3, 1), NewVolumeRandom(3, 3, 3), false}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if got := vol.EqualSize(tt.vol2); got != tt.want {
				t.Errorf("Volume.EqualSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_MulElem2(t *testing.T) {
	tests := []struct {
		name string
		vol1 Volume
		vol2 Volume
		want Volume
	}{
		{"All Zeros", NewVolume(3, 3, 3), NewVolume(3, 3, 3), NewVolume(3, 3, 3)},
		{"Zeros with random", NewVolume(3, 3, 3), NewVolumeRandom(3, 3, 3), NewVolume(3, 3, 3)},
		{"Random with zeros", NewVolumeRandom(3, 3, 3), NewVolume(3, 3, 3), NewVolume(3, 3, 3)},
		{"testVol squared", testVol, testVol, testVolSquared}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			vol := tt.vol1
			if vol.MulElem2(tt.vol2); !vol.Equals(tt.want) {
				t.Errorf("Volume.MulElem2 = %v, want %v", vol, tt.want)
			}
		})
	}
}

func TestVolume_Max(t *testing.T) {

	tests := []struct {
		name string
		vol  Volume
		want float64
	}{
		{"Numbered", testVol, 26},
		{"All Zero", NewVolume(5, 5, 5), 0}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.vol.Max(); got != tt.want {
				t.Errorf("Volume.Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVolume_SimimlarTo(t *testing.T) {
	type fields struct {
		Fields []mat64.Dense
	}
	type args struct {
		in        Volume
		threshold float64
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
			vol := Volume{
				Fields: tt.fields.Fields,
			}
			if got := vol.SimimlarTo(tt.args.in, tt.args.threshold); got != tt.want {
				t.Errorf("Volume.SimimlarTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
