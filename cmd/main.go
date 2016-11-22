package main

import (
	"github.com/binaryplease/robonet"
)

func main() {

	/*

	   Example: Blur Filter


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

	   	//Add a poolig layer to reduce image size by half
	   	layPool := new(robonet.PoolLayer)
	   	layPool.StrideR = 2
	   	layPool.StrideC = 2
	   	layPool.SizeR = 2
	   	layPool.SizeC = 2
	   	net.AddLayer(layPool)

	   	//Set net's input
	   	net.Input = robonet.VolumeFromTIFF("images/bw5.tiff")

	   	//Calculate the net's outpout
	   	net.Calculate()

	   	//Save the output to B/W image
	   	robonet.SaveVolumeToTIFF("out.tiff", net.Output)
	*/

	//Example: Edge detection

	//Create Net
	net := new(robonet.Net)

	//Create ConvLayer
	layConv := new(robonet.ConvLayer)

	//Create kernel filled with -1 and 8 (edge detection)
	imgker1 := robonet.NewKernelFilled(3, 3, 3, -1)

	imgker1.SetAt(1, 1, 0, 8)
	imgker1.SetAt(1, 1, 1, 8)
	imgker1.SetAt(1, 1, 2, 8)

	//Add kernel to ConvLayer
	layConv.AddKernel(imgker1, 1, 1)

	//Add ConvLayer to net
	net.AddLayer(layConv)

	//Add normalisation Layer to limit output to 255 max
	layNorm1 := new(robonet.NormLayer)
	layNorm1.NormVal = 255
	net.AddLayer(layNorm1)

	//Set net's input
	net.Input = robonet.VolumeFromTIFF("images/bwcircle.tiff")

	//Calculate the net's outpout
	net.Calculate()

	//Save the output to B/W image
	robonet.SaveVolumeToTIFF("out.tiff", net.Output)
}
