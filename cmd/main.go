package main

import (
	"fmt"
	"gitlab.com/binaryplease/robonet"
)

func main() {
	//Volume and Kernel

	inputVol := robonet.VolumeFromImageFile("images/test.jpg")

	net := new(robonet.Net)

	fmt.Println("Setting input")
	inputVol.Print()
	net.Input = inputVol

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

	net.AddLayer(new(robonet.InputLayer))
	net.AddLayer(new(robonet.ConvLayer))
	net.AddLayer(new(robonet.PoolLayer))
	net.AddLayer(new(robonet.NormLayer))
	net.AddLayer(new(robonet.FCLayer))
	net.AddLayer(new(robonet.ReluLayer))

	net.AddLayer(layIn)
	net.AddLayer(layConv)
	net.AddLayer(new(robonet.PoolLayer))

	fmt.Println("calculate output")
	net.Calculate()

	fmt.Println("output was")
	//net.Output.Print()
	robonet.SaveVolumeToFile("out.jpg", *robonet.NewVolumeRandom(300, 300, 3))

}
