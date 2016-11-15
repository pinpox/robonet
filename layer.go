package robonet

//RNConvLayer basic type for a convolutional layer
type RNConvLayer struct {
	kernels  []Kernel
	stridesR []int
	stridesC []int
}

//AddKernel adds a kernel to a layer
func (l *RNConvLayer) AddKernel(kern Kernel, strideR, strideC int) {
	l.kernels = append(l.kernels, kern)
	l.stridesR = append(l.stridesR, strideR)
	l.stridesC = append(l.stridesC, strideC)
}

//Calculate applys all Kernels to a given Volume
func (l *RNConvLayer) Calculate(vol Volume) Volume {
	//result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	for i, v := range l.kernels {
		vol.Apply(v, l.stridesR[i], l.stridesC[i])
	}
	return vol
}

//Kernels returns the kernels of the layer
func (l RNConvLayer) Kernels() []Kernel {
	return l.kernels

}

// Layer represents the general type of all layer types
type Layer interface {
}
