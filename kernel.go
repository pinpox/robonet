package robonet

import (
	"fmt"
	_ "github.com/gonum/matrix/mat64"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	values Volume
}

func (kern *Kernel) GetAt(r, c, d int) float64 {
	return kern.values.GetAt(r, c, d)
}

func (kern *Kernel) SetAt(r, c, d int, val float64) {
	kern.values.SetAt(r, c, d, val)
}

func (kern *Kernel) SetAll(v Volume) {

	r, c, d := kern.Dims()
	if !EqualVolDim(kern.Vol(), v) {
		panic("Volumedimensions do not match!")
	}
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < d; k++ {
				kern.values.SetAt(i, j, k, v.GetAt(i, j, k))
			}
		}
	}
}

//Vol returns the underlying volume of a kernel
func (kern *Kernel) Vol() Volume {
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

//Equals compares to kernels
func (kern *Kernel) Equals(in Kernel) bool {
	return kern.values.Equals(in.values)
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
func (kern Kernel) Print() {
	kern.values.Print()
}

//Dims returns the  size of the kernel
func (kern Kernel) Dims() (int, int, int) {
	return kern.values.Dims()
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

func (kern *Kernel) PointReflect() {
	kern.values.PointReflect()
}

func (kern *Kernel) Reflect() {
	kern.values.Reflect()
}
