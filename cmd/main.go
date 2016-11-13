package main

import (
	"fmt"
	"gitlab.com/binaryplease/robonet"
)

func main() {
	//Volume and Kernel

	inputVol := *robonet.NewRNVolumeRandom(4, 4, 3)

	fmt.Println("input was")
	inputVol.Print()

	fmt.Println("Create a new Layer")
	lay := new(robonet.RNConvLayer)

	fmt.Println("add a kernel 1")
	kernel1 := robonet.NewKernelRandom(3, 3, 3)
	kernel1.Print()
	lay.AddKernel(*kernel1)

	fmt.Println("add a kernel 2")
	kernel2 := robonet.NewKernelRandom(3, 3, 2)
	kernel2.Print()
	lay.AddKernel(*kernel2)

	fmt.Println("calculate output")
	outputVol := lay.Calculate(inputVol)

	fmt.Println("output was")
	outputVol.Print()
}
