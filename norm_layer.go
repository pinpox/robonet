package robonet

//NormLayer is a normalisation layer
type NormLayer struct {
	LayerFields
}

//Calculate applyes the normalisation funktion for every element in the input volume
func (lay *NormLayer) Calculate() {
	lay.output = lay.input
	for r := 0; r < lay.input.Rows(); r++ {
		for c := 0; c < lay.input.Collumns(); c++ {
			for d := 0; d < lay.input.Depth(); d++ {
				lay.output.SetAt(r, c, d, SigmoidFast(lay.input.GetAt(r, c, d)))
			}
		}
	}
}
