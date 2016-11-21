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
	imgker1 := robonet.NewKernelFilled(9, 9, 3, 1)
	imgker2 := robonet.NewKernelFilled(9, 9, 3, 1)
	imgker3 := robonet.NewKernelFilled(9, 9, 3, 1)

	//Add kernel to ConvLayer
	layConv.AddKernel(imgker1, 1, 1)
	layConv.AddKernel(imgker2, 1, 1)
	layConv.AddKernel(imgker3, 1, 1)

	//Add ConvLayer to net
	net.AddLayer(layConv)

	//Add normalisation Layer to limit output to 255 max
	layNorm1 := new(robonet.NormLayer)
	layNorm1.NormVal = 255
	net.AddLayer(layNorm1)

	//Set net's input
	net.Input = robonet.VolumeFromTIFF("images/bw5.tiff")

	//Calculate the net's outpout
	net.Calculate()

	//Save the output to B/W image
	robonet.SaveVolumeToTIFF("out.tiff", net.Output)

}
