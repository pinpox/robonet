package robonet

import (
	"fmt"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	values rNVolume
}

//NewKernel creates a new kernel initialized with zeros
func NewKernel(height int, width, depth int) *Kernel {
	if !Odd3Dim(height, width, depth) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewRNVolume(height, width, depth)}
	return &g
}

//NewKernelRandom creates a new kernel initialized with random values
func NewKernelRandom(height int, width, depth int) *Kernel {
	if Odd3Dim(height, width, depth) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewRNVolumeRandom(height, width, depth)}
	return &g
}

//Print shows show the kernel's matrix string representation
func (f Kernel) Print() {
	f.values.Print()
}

//Dims returns the  size of the kernel
func (f Kernel) Dims() (int, int, int) {
	return f.values.Dims()
}

//Apply applys the kernel to a equally sized chunk of a volume
//Only kernels of the same size as the volume can be applied
func (f Kernel) Apply(in rNVolume) float64 {

	ConvResult := 1.0

	if !(f.values.EqualSize(in)) {
		fmt.Println("Kernel size doesn't match input")
		panic("Kernel size doesn't match input")
	}

	// 1) reflect kernel
	// 2) multiply pairwise
	// 3) normalize

	return ConvResult
}

func (f Kernel) pointReflection(input Kernel) Kernel {
	//TODO
	return input
}

func (f Kernel) reflection(input Kernel) Kernel {
	//TODO
	return input
}
