package robonet

//RNConvLayer basic type for a convolutional layer
type RNConvLayer struct {
	Kernels []Kernel
}

//AddKernel adds a kernel to a layer
func (l RNConvLayer) AddKernel(fil *Kernel) {
	l.Kernels = append(l.Kernels, *fil)
}

//Calculate applys all Kernels to a given Volume
func (l RNConvLayer) Calculate(vol rNVolume) rNVolume {
	//result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	for _, v := range l.Kernels {
		vol = vol.Apply(v)
	}
	return vol
}

// Layer represents the general type of all layer types
type Layer interface {
}
