package main

import (
	"fmt"
	"gitlab.com/binaryplease/robonet"
)

func main() {

	inputVol := *robonet.NewRNVolumeRandom(4, 4, 3)
	//rNVolume{
	//[]mat64.Dense{
	//*mat64.NewDense(4, 4, nil),
	//*mat64.NewDense(4, 4, nil),
	//*mat64.NewDense(4, 4, nil)}}

	fmt.Println("input was")
	inputVol.Print()

	fmt.Println("Create a new Layer")
	lay := new(robonet.RNConvLayer)

	fmt.Println("add a filter 1")
	filter1 := robonet.NewFilterRandom(3, 3, 3)
	filter1.Print()
	lay.AddFilter(filter1)

	fmt.Println("add a filter 2")
	filter2 := robonet.NewFilterRandom(2, 4, 2)
	filter2.Print()
	lay.AddFilter(filter2)

	fmt.Println("calculate output")
	outputVol := lay.Calculate(inputVol)

	fmt.Println("output was")
	outputVol.Print()
}