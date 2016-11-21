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
		fmt.Println("Started layer: ", k)
		fmt.Printf("Input has Dims %vx%vx%v \n", res.Rows(), res.Collumns(), res.Depth())
		v.Input(res)
		fmt.Println("Calculating: ", k)
		v.Calculate()
		res = v.Output()
		fmt.Printf("output has Dims %vx%vx%v\n", res.Rows(), res.Collumns(), res.Depth())
	}

	fmt.Printf("totoal output has dims %vx%vx%v\n", res.Rows(), res.Collumns(), res.Depth())
	//res.Print()
	net.Output = res
}
