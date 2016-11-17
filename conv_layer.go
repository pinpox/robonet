package robonet

//ConvLayer basic type for a convolutional layer
//The layer will compute the output of neurons that are connected to local regions in the input,
//each computing a dot product between their weights and a small region they are connected to in
//the input volume. This may result in volume such as [32x32x12] if we decided to use 12 filters.
type ConvLayer struct {
	LayerFields
	kernels  []Kernel
	stridesR []int
	stridesC []int
}

//AddKernel adds a kernel to a layer
func (l *ConvLayer) AddKernel(kern Kernel, strideR, strideC int) {
	l.kernels = append(l.kernels, kern)
	l.stridesR = append(l.stridesR, strideR)
	l.stridesC = append(l.stridesC, strideC)
}

//Calculate applys all Kernels to a given Volume
func (l *ConvLayer) Calculate() {

	l.output = *NewVolume(10, 10, len(l.kernels)) //TODO
	for k, v := range l.kernels {

		for r := 0; r < l.input.Rows(); r++ {
			for c := 0; c < l.input.Collumns(); c++ {

				l.output.SetAt(r, c, k, v.Apply(l.input.SubVolumePadded(r*l.stridesR[k], c*l.stridesC[k], v.Rows(), v.Collumns())))
			}
		}
	}
}

//Kernels returns the kernels of the layer
func (l ConvLayer) Kernels() []Kernel {
	return l.kernels
}
