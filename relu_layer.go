package robonet

//ReluLayer will apply an elementwise activation function, such as the max(0,x)max(0,x)
//thresholding at zero. This leaves the size of the volume unchanged ([32x32x12]).
type ReluLayer struct {
	LayerFields
}

//Calculate for ReluLayer
func (lay *ReluLayer) Calculate() {
	lay.output = lay.input
	for r := 0; r < lay.input.Rows(); r++ {
		for c := 0; c < lay.input.Collumns(); c++ {
			for d := 0; d < lay.input.Depth(); d++ {
				lay.output.SetAt(r, c, d, SigmoidFast(lay.input.GetAt(r, c, d)))
			}
		}
	}
}
