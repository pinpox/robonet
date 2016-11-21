package robonet

import (
	"errors"
	"fmt"
	"log"
)

//ConvLayer basic type for a convolutional layer
//The layer will compute the output of neurons that are connected to local regions in the input,
//each computing a dot product between their weights and a small region they are connected to in
//the input volume. This may result in volume such as [32x32x12] if we decided to use 12 filters.
type ConvLayer struct {
	LayerFields
	kernels []Kernel
	strideR int
	strideC int
}

//AddKernel adds a kernel to a layer
func (l *ConvLayer) AddKernel(kern Kernel, strideR, strideC int) {
	l.kernels = append(l.kernels, kern)

	if l.strideC == 0 || l.strideR == 0 {
		if strideC == 0 || strideR == 0 {
			log.Fatal(errors.New("robonet.ConvLayer: invalid stride for kernel"))
		}
		l.strideC = strideC
		l.strideR = strideR

	} else {

		if l.strideC != strideC || l.strideR != strideR {
			log.Fatal(errors.New("robonet.ConvLayer: invalid stride for kernel, already set"))

		}
	}
}

//Calculate applys all Kernels to a given Volume
func (l *ConvLayer) Calculate() {

	l.output = NewVolume(l.input.Rows()/l.strideR, l.input.Collumns()/l.strideC, len(l.kernels))
	for k, v := range l.kernels {

		fmt.Println("	Calculating Kernel ", k)
		for r := 0; r < l.output.Rows(); r++ {
			for c := 0; c < l.output.Collumns(); c++ {

				l.output.SetAt(r, c, k, v.Apply(l.input.SubVolumePadded(r*l.strideR, c*l.strideC, v.Rows(), v.Collumns())))
			}
		}
	}
}

//Kernels returns the kernels of the layer
func (l ConvLayer) Kernels() []Kernel {
	return l.kernels
}
