package robonet

import "testing"

func TestNet_AddLayer(t *testing.T) {
	//trivial
}

func TestNet_Calculate(t *testing.T) {
	type fields struct {
		layers []Layer
		Input  Volume
		Output Volume
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			net := &Net{
				layers: tt.fields.layers,
				Input:  tt.fields.Input,
				Output: tt.fields.Output,
			}
			net.Calculate()
		})
	}
}
