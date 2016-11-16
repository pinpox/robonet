package robonet

//FCLayer (i.e. fully-connected) layer will compute the class scores, resulting in volume of size [1x1x10],
//where each of the 10 numbers correspond to a class score, such as among the 10 categories of CIFAR-10.
//As with ordinary Neural Networks and as the name implies, each neuron in this layer will be connected to all the numbers in the previous volume.
type FCLayer struct {
	LayerFields
	//TODO
}

//Calculate method for fully connected layers
func (lay *FCLayer) Calculate() {
	//TODO
}
