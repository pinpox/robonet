package robonet

//Net is the basic type for Conv nets
type Net struct {
	layers []Layer
	Input  Volume
	Output Volume
}

//AddLayer adds another layer to the net
func (net *Net) AddLayer(lay Layer) {
	net.layers = append(net.layers, lay)
}

//Calculate calcuates te Output
func (net *Net) Calculate() {

	res := net.Input

	for _, v := range net.layers {
		v.Input(res)
		v.Calculate()
		res = v.Output()
	}

	net.Output = res
}
