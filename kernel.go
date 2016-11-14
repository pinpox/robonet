package robonet

import (
	"fmt"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	values rNVolume
}

func (kern *Kernel) GetAt(r, c, d int) float64 {
	return kern.values.GetAt(r,c,d)
}

func (kern *Kernel) SetAt(r, c, d int, val float64) {
	kern.values.SetAt(r, c, d, val)
}


func (kern *Kernel) SetAll(v rNVolume) {

	r,c,d := kern.Dims()
	if !EqualVolDim(kern.Vol(), v) {
			panic("Volumedimensions do not match!")
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k:= 0; k < d; k++ {
				kern.values.SetAt(i,j,k, v.GetAt(i,j,k))
			}
		}
	}
}

func  (kern *Kernel) Vol() rNVolume {
	return kern.values
}


//NewKernel creates a new kernel initialized with zeros
func NewKernel(r, c, d int) Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewRNVolume(r, c, d)}
	return g
}

//NewKernelRandom creates a new kernel initialized with random values
func NewKernelRandom(r, c, d int) *Kernel {
	if !Odd3Dim(r, c, d) {
		panic("Kernel must have odd width and heigth")
	}
	g := Kernel{*NewRNVolumeRandom(r, c, d)}
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





func (f *Kernel) PointReflect() {
	f.values.PointReflect()
}

func (f *Kernel) Reflect() {
	f.values.Reflect()
}
