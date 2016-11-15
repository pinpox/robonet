package main

import (
	"fmt"
	"gitlab.com/binaryplease/robonet"
)

func main() {
	//Volume and Kernel

	inputVol := *robonet.NewVolumeRandom(4, 4, 3)

	//fmt.Println("input was")
	//inputVol.Print()

	net := new(robonet.Net)

	fmt.Println("Create a new Layer")
	layConv := new(robonet.ConvLayer)

	fmt.Println("add a kernel 1")
	kernel1 := robonet.NewKernelRandom(3, 3, 3)
	kernel1.Print()
	layConv.AddKernel(*kernel1, 1, 1)

	fmt.Println("add a kernel 2")
	kernel2 := robonet.NewKernelRandom(3, 3, 2)
	kernel2.Print()
	layConv.AddKernel(*kernel2, 1, 1)

	layIn := new(robonet.InputLayer)
	layIn.SetInput(inputVol)

	net.AddLayer(layIn)
	net.AddLayer(layConv)
	net.AddLayer(new(robonet.PoolLayer))

	fmt.Println("calculate output")
	out := net.Calculate()

	fmt.Println("output was")
	out.Print()

}
