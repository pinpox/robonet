package main

import (
	"gitlab.com/binaryplease/robonet"
)

func main() {

	//Create Net
	net := new(robonet.Net)

	//Create ConvLayer
	layConv := new(robonet.ConvLayer)

	//Create kernel filled with 1s (blur)
	imgker := robonet.NewKernelFilled(9, 9, 3, 1)

	//Add kernel to ConvLayer
	layConv.AddKernel(imgker, 1, 1)

	//Add ConvLayer to net
	net.AddLayer(layConv)

	//Set net's input
	net.Input = robonet.VolumeFromTIFF("images/bw5.tiff")

	//Calculate the net's outpout
	net.Calculate()

	//Normalize output to 255 max
	net.Output.Norm(255.0)

	//Save the output to B/W image
	robonet.SaveVolumeToTIFF("out.tiff", net.Output)

}
