package robonet

//ReluLayer will apply an elementwise activation function, such as the max(0,x)max(0,x)
//thresholding at zero. This leaves the size of the volume unchanged ([32x32x12]).
type ReluLayer struct {
	LayerFields
	//TODO
}

//Calculate for ReluLayer
func (lay *ReluLayer) Calculate() {
	//TODO
}
