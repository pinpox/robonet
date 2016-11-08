package main

import (
	"fmt"
	//"github.com/gonum/matrix/mat64"
)

func main() {

	inputVol := *newRNVolumeRandom(4, 4, 3)
	//rNVolume{
	//[]mat64.Dense{
	//*mat64.NewDense(4, 4, nil),
	//*mat64.NewDense(4, 4, nil),
	//*mat64.NewDense(4, 4, nil)}}

	fmt.Println("input was")
	inputVol.Print()

	fmt.Println("Create a new Layer")
	lay := new(rNConvLayer)

	fmt.Println("add a filter 1")
	filter1 := NewFilterRandom(3, 3, 3)
	filter1.Print()
	lay.AddFilter(filter1)

	fmt.Println("add a filter 2")
	filter2 := NewFilterRandom(2, 4, 2)
	filter2.Print()
	lay.AddFilter(filter2)

	fmt.Println("calculate output")
	outputVol := lay.calculate(inputVol)

	fmt.Println("output was")
	outputVol.Print()
}
