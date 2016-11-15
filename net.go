package robonet

//Net is the basic type for Conv nets
type Net struct {
	layers []Layer
}

//AddLayer adds another layer to the net
func (net *Net) AddLayer(lay Layer) {
	net.layers = append(net.layers, lay)
}

//Calculate calcuates te output
func (net Net) Calculate() Volume {
	//TODO
	return *NewVolumeRandom(3, 3, 3)

}
