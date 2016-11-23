package robonet

import (
	"reflect"
	"testing"
)

func Test_maxPool(t *testing.T) {
	tests := []struct {
		name    string
		vol     Volume
		wantRes []float64
	}{
		{"All Zeros", New(10, 10, 5), []float64{0, 0, 0, 0, 0}},
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

	res1 := New(2, 2, 3)

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

	res2 := New(1, 1, 3)
	res2.SetAt(0, 0, 0, 8)
	res2.SetAt(0, 0, 1, 17)
	res2.SetAt(0, 0, 2, 26)

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
		{"All Zeros stride 2 size 2", New(6, 6, 3), fields{2, 2, 2, 2}, New(3, 3, 3)},
		{"All Zeros stride 3 size 3", New(6, 6, 3), fields{3, 3, 3, 3}, New(2, 2, 3)},
		{"All Zeros stride 6 size 6", New(6, 6, 3), fields{6, 6, 6, 6}, New(1, 1, 3)},
		{"All Zeros stride 2 size 4", New(6, 6, 3), fields{4, 4, 2, 2}, New(2, 2, 3)},
		{"All Zeros stride 6 size 4", New(10, 10, 3), fields{4, 4, 6, 6}, New(2, 2, 3)},
		{"All Zeros stride 5 size 2", New(10, 10, 3), fields{2, 2, 5, 5}, New(2, 2, 3)},
		{"testVol stride 5 size 2", testVol, fields{2, 2, 1, 1}, res1},
		{"testVol stride 5 size 2", testVol, fields{3, 3, 1, 1}, res2},
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
