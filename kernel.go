package robonet

import (
	"errors"
	"fmt"
	"log"
)

// Kernel represets a basic conv kernel
type Kernel struct {
	Volume
}

//NewKernel creates a new kernel initialized with zeros
func NewKernel(r, c, d int) Kernel {
	if !Odd3Dim(r, c, d) {
		log.Fatal(errors.New("Kernel must have odd width and heigth"))
	}
	g := Kernel{New(r, c, d)}
	return g
}

//Equals compares to kernels
func (kern *Kernel) Equals(in Kernel) bool {
	return kern.Volume.Equals(in.Volume)
}

//NewKernelRandom creates a new kernel initialized with random values
func NewKernelRandom(r, c, d int) Kernel {
	if !Odd3Dim(r, c, d) {
		log.Fatal(errors.New("Kernel must have odd width and heigth"))
	}
	g := Kernel{NewRand(r, c, d)}
	return g
}

//NewKernelFilled creates a new kernel initialized with random values
func NewKernelFilled(r, c, d int, fil float64) Kernel {
	if !Odd3Dim(r, c, d) {
		log.Fatal(errors.New("Kernel must have odd width and heigth"))
	}
	g := Kernel{NewFull(r, c, d, fil)}
	return g
}

//Apply applys the kernel to a equally sized chunk of a volume
//Only kernels of the same size as the volume can be applied
func (kern Kernel) Apply(in Volume) float64 {

	ConvResult := 0.0
	s := kern.Shape()

	if !(kern.Volume.EqualSize(in)) {
		if in == nil {

			panic("nil")
			fmt.Printf("kernel: %vx%vx%v vol: nil", s[0], s[1], s[2])
			log.Fatal(errors.New("Kernel size doesn't match input "))
		}
		fmt.Printf("kernel: %vx%vx%v, vol: %vx%vx%v", s[0], s[1], s[2], in.Rows(), in.Collumns(), in.Depth())
		log.Fatal(errors.New("Kernel size doesn't match input "))
	}

	// 1) reflect kernel
	kernRef := kern
	kernRef.PointReflect()
	// 2) multiply pairwise

	res := New(s[0], s[1], s[2])

	res.SetAll(kern.Volume)

	res.MulElem(in)

	for i := 0; i < s[0]; i++ {
		for j := 0; j < s[1]; j++ {
			for k := 0; k < s[2]; k++ {
				// TODO check if normalization is needed!
				ConvResult += (res.GetAt(i, j, k))
			}
		}
	}

	return ConvResult
}

//Elems returns the number of element a kernel has
func (kern Kernel) Elems() int {
	return kern.Volume.Elems()
}

//Sum returns the sum of all elements in the kernel
func (kern Kernel) Sum() float64 {
	res := 0.0
	for r := 0; r < kern.Rows(); r++ {
		for c := 0; c < kern.Collumns(); c++ {
			for d := 0; d < kern.Depth(); d++ {
				res += kern.GetAt(r, c, d)
			}
		}
	}

	return res

}
