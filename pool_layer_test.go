package robonet

import (
	"reflect"
	"testing"
)

func TestPoolLayer_Calculate(t *testing.T) {
	type fields struct {
		LayerFields LayerFields
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lay := &PoolLayer{
				LayerFields: tt.fields.LayerFields,
			}
			lay.Calculate()
		})
	}
}

func Test_maxPool(t *testing.T) {
	type args struct {
		vol Volume
	}
	tests := []struct {
		name    string
		args    args
		wantRes []float64
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := maxPool(tt.args.vol); !reflect.DeepEqual(gotRes, tt.wantRes) {
				t.Errorf("maxPool() = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}
