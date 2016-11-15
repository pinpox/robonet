package robonet

import (
	"fmt"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	Volume
}

//GetAt returns the value of the volume of the kernel at a given position
func (kern *Kernel) GetAt(r, c, d int) float64 {
	return kern.GetAt(r, c, d)
}

//SetAt sets the value of the volume of the kernel at a given position
func (kern *Kernel) SetAt(r, c, d int, val float64) {
	kern.SetAt(r, c, d, val)
}

//SetAll sets all values of the kernel's volume from another equal-sized volume
func (kern *Kernel) SetAll(v Volume) {

	r, c, d := kern.Dims()
	if !EqualVolDim(kern.Volume, v) {
		panic("Volumedimensions do not match!")
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < d; k++ {
				kern.SetAt(i, j, k, v.GetAt(i, j, k))
			}
		}
	}
}

//NewKernel creates a new kernel initialized with zeros
func NewKernel(r, c, d int) Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewRNVolume(r, c, d)}
	return g
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
	g := Kernel{*NewRNVolumeRandom(r, c, d)}
	return &g
}

//Apply applys the kernel to a equally sized chunk of a volume
//Only kernels of the same size as the volume can be applied
func (kern Kernel) Apply(in Volume) float64 {

	ConvResult := 0.0
	r, c, d := kern.Dims()

	if !(kern.values.EqualSize(in)) {
		fmt.Println("Kernel size doesn't match input")
		panic("Kernel size doesn't match input")
	}

	// 1) reflect kernel
	kernRef := kern
	kernRef.PointReflect()
	// 2) multiply pairwise

	res := NewRNVolume(r, c, d)

	*res = kern.values

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

//PointReflect calculates the pointreflection of the kernel's volume
func (kern *Kernel) PointReflect() {
	kern.values.PointReflect()
}

//Reflect calculates the reflection of the kernel's volume
func (kern *Kernel) Reflect() {
	kern.values.Reflect()
}
