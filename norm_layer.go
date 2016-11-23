package robonet

//NormLayer is a normalisation layer
type NormLayer struct {
	LayerFields
	NormVal float64
}

//Calculate applyes the normalisation funktion for every element in the input volume
func (lay *NormLayer) Calculate() {
	lay.output = lay.input
	lay.output.Norm(lay.NormVal)
}
