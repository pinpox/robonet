package robonet

// Layer interface
type Layer interface {
	Input(Volume)
	Calculate()
	Output() Volume
}

//LayerFields basic data fields every layertype should have
type LayerFields struct {
	input  Volume
	output Volume
}

// Input is the Default method for Setting the input of a layer
func (lf *LayerFields) Input(vol Volume) {
	lf.input = vol
}

// Output is the default method for retrieving the output of a layer
func (lf *LayerFields) Output() Volume {
	return lf.output
}
