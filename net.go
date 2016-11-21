package robonet

import (
	"fmt"
)

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

	for k, v := range net.layers {
		fmt.Printf("[Layer-%v] Start\n", k)

		fmt.Printf("[Layer-%v] Input Dims %vx%vx%v \n", k, res.Rows(), res.Collumns(), res.Depth())
		v.Input(res)

		fmt.Printf("[Layer-%v] Calculating\n", k)
		v.Calculate()

		res = v.Output()
		fmt.Printf("[Layer-%v] Output Dims %vx%vx%v\n", k, res.Rows(), res.Collumns(), res.Depth())
	}

	fmt.Printf("Total output Dims %vx%vx%v\n", res.Rows(), res.Collumns(), res.Depth())
	net.Output = res
}
