package robonet

//ConvLayer basic type for a convolutional layer
//The layer will compute the output of neurons that are connected to local regions in the input,
//each computing a dot product between their weights and a small region they are connected to in
//the input volume. This may result in volume such as [32x32x12] if we decided to use 12 filters.
type ConvLayer struct {
	Layer
	kernels  []Kernel
	stridesR []int
	stridesC []int
}

//PoolLayer will perform a downsampling operation along the spatial dimensions (width, height), resulting in volume such as [16x16x12].
type PoolLayer struct {
	Layer
	//TODO
}

//NormLayer is a normalisation layer
type NormLayer struct {
	Layer
	//TODO
}

//FCLayer (i.e. fully-connected) layer will compute the class scores, resulting in volume of size [1x1x10], where each of the 10 numbers correspond to a class score, such as among the 10 categories of CIFAR-10. As with ordinary Neural Networks and as the name implies, each neuron in this layer will be connected to all the numbers in the previous volume.
type FCLayer struct {
	Layer
	//TODO
}

//InputLayer [32x32x3] will hold the raw pixel values of the image, in this case an image of width 32, height 32, and with three color channels R,G,B.
type InputLayer struct {
	Layer
	input Volume
	//TODO
}

//SetInput sets the input of a input layer
func (l *InputLayer) SetInput(in Volume) {
	l.input = in
}

//ReluLayer will apply an elementwise activation function, such as the max(0,x)max(0,x)
//thresholding at zero. This leaves the size of the volume unchanged ([32x32x12]).
type ReluLayer struct {
	Layer
	//TODO
}

//AddKernel adds a kernel to a layer
func (l *ConvLayer) AddKernel(kern Kernel, strideR, strideC int) {
	l.kernels = append(l.kernels, kern)
	l.stridesR = append(l.stridesR, strideR)
	l.stridesC = append(l.stridesC, strideC)
}

//Calculate applys all Kernels to a given Volume
func (l *ConvLayer) Calculate(vol Volume) Volume {
	//TODO
	//result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	//for i, v := range l.kernels {
	//vol.Apply(v, l.stridesR[i], l.stridesC[i])
	//}
	return vol
}

//Kernels returns the kernels of the layer
func (l ConvLayer) Kernels() []Kernel {
	return l.kernels

}

// Layer represents the general type of all layer types
type Layer interface {
}
