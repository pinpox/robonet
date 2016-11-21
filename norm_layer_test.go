package robonet

import "testing"

func TestNormLayer_Calculate(t *testing.T) {
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
			lay := &NormLayer{
				LayerFields: tt.fields.LayerFields,
			}
			lay.Calculate()
		})
	}
}
