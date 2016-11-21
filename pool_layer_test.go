package robonet

import (
	//"github.com/gonum/matrix/mat64"
	"reflect"
	"testing"
)

//var data1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
//var data2 = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
//var data3 = []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
//var testVol = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, data1), *mat64.NewDense(3, 3, data2), *mat64.NewDense(3, 3, data3)}}

func Test_maxPool(t *testing.T) {
	tests := []struct {
		name    string
		vol     Volume
		wantRes []float64
	}{
		{"All Zeros", NewVolume(10, 10, 5), []float64{0, 0, 0, 0, 0}},
		{"testVol", testVol, []float64{8, 17, 26}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxPool(tt.vol); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("maxPool() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestPoolLayer_Calculate(t *testing.T) {

	res1 := NewVolume(2, 2, 3)

	res1.SetAt(0, 0, 0, 4)
	res1.SetAt(0, 1, 0, 5)
	res1.SetAt(1, 0, 0, 7)
	res1.SetAt(1, 1, 0, 8)

	res1.SetAt(0, 0, 1, 13)
	res1.SetAt(0, 1, 1, 14)
	res1.SetAt(1, 0, 1, 16)
	res1.SetAt(1, 1, 1, 17)

	res1.SetAt(0, 0, 2, 22)
	res1.SetAt(0, 1, 2, 23)
	res1.SetAt(1, 0, 2, 25)
	res1.SetAt(1, 1, 2, 26)

	type fields struct {
		SizeR   int
		SizeC   int
		StrideR int
		StrideC int
	}
	tests := []struct {
		name   string
		vol    Volume
		fields fields
		want   Volume
	}{
		{"All Zeros stride 2 size 2", NewVolume(6, 6, 3), fields{2, 2, 2, 2}, NewVolume(3, 3, 3)},
		{"All Zeros stride 3 size 3", NewVolume(6, 6, 3), fields{3, 3, 3, 3}, NewVolume(2, 2, 3)},
		{"All Zeros stride 6 size 6", NewVolume(6, 6, 3), fields{6, 6, 6, 6}, NewVolume(1, 1, 3)},
		{"All Zeros stride 2 size 4", NewVolume(6, 6, 3), fields{4, 4, 2, 2}, NewVolume(2, 2, 3)},
		{"All Zeros stride 6 size 4", NewVolume(10, 10, 3), fields{4, 4, 6, 6}, NewVolume(2, 2, 3)},
		{"All Zeros stride 5 size 2", NewVolume(10, 10, 3), fields{2, 2, 5, 5}, NewVolume(2, 2, 3)},
		{"testVol stride 5 size 2", testVol, fields{2, 2, 1, 1}, res1},

		//TODO add test cases wit numbers
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lay := &PoolLayer{
				SizeR:   tt.fields.SizeR,
				SizeC:   tt.fields.SizeC,
				StrideR: tt.fields.StrideR,
				StrideC: tt.fields.StrideC,
			}

			lay.Input(tt.vol)
			lay.Calculate()

			if got := lay.Output(); !got.Equals(tt.want) {
				t.Errorf("PoolLayer.Calculate() = %v, want %v", got, tt.want)
			}
		})
	}
}
