package robonet

import (
	"fmt"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	Volume
}

//NewKernel creates a new kernel initialized with zeros
func NewKernel(r, c, d int) *Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewVolume(r, c, d)}
	return &g
}

//Equals compares to kernels
func (kern *Kernel) Equals(in Kernel) bool {
	return kern.Volume.Equals(in.Volume)
}

//NewKernelRandom creates a new kernel initialized with random values
func NewKernelRandom(r, c, d int) *Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewVolumeRandom(r, c, d)}
	return &g
}

//NewKernelFilled creates a new kernel initialized with random values
func NewKernelFilled(r, c, d int, fil float64) *Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewVolumeFilled(r, c, d, fil)}
	return &g
}

//Apply applys the kernel to a equally sized chunk of a volume
//Only kernels of the same size as the volume can be applied
func (kern Kernel) Apply(in Volume) float64 {

	ConvResult := 0.0
	r, c, d := kern.Dims()

	if !(kern.Volume.EqualSize(in)) {
		fmt.Println("Kernel size doesn't match input")
		panic("Kernel size doesn't match input")
	}

	// 1) reflect kernel
	kernRef := kern
	kernRef.PointReflect()
	// 2) multiply pairwise

	res := NewVolume(r, c, d)

	*res = kern.Volume

	res.MulElem2(in)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < d; k++ {
				ConvResult += res.GetAt(i, j, k)
			}
		}
	}

	// 3) normalize??

	return ConvResult
}
