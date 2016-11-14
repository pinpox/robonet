package robonet

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
)

// rNVolume is a basic type to hold the layer's information
type rNVolume struct {
	Fields []mat64.Dense
}

func (vol *rNVolume) Dims() (int, int, int) {
	r, c := vol.Fields[0].Dims()
	d := len(vol.Fields)
	return r, c, d
}

//Apply applys the given kernel to the whole volume, returnung a Volume with 1 depth
func (vol *rNVolume) Apply(f Kernel) rNVolume {

	//TODO apply the kernel to the volume

	//Check correct output
	_, _, a := vol.Dims()
	if a != 1 {
		panic("should have returned a plane (2dim)")
	}

	return *vol
}

//NewRNVolume generates a rNVolume of fixed size filled with zeros
func NewRNVolume(r, c, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *mat64.NewDense(r, c, nil))
	}
	return v
}

//NewRNVolumeRandom generates a rNVolume of fixed size filled with values between 0 and 1
func NewRNVolumeRandom(r, c, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	for j := 0; j < d; j++ {

		data := make([]float64, r*c)
		for i := range data {
			data[i] = rand.Float64()
		}
		a := mat64.NewDense(r, c, data)

		v.Fields = append(v.Fields, *a)
	}
	return v
}

//SubVolumePadded returns a part of the original Volume. cR and cC determine the center of copying, r and c the size of the subvolume.
//If the size exceeds the underlying volume the submodule is filled(padded with Zeros.
func (vol *rNVolume) SubVolumePadded(cR, cC, r, c int) rNVolume {

	if r%2 == 0 || c%2 == 0 {
		panic("Even dimensions not allowed for subvolumes")
	}

	sub := NewRNVolume(r, c, vol.Depth())

	for ir := 0; ir < sub.Rows(); ir++ {
		for ic := 0; ic < sub.Collumns(); ic++ {
			for id := 0; id < sub.Depth(); id++ {

				val := 0.0

				offsetR := ((vol.Rows() - 1) / 2) - cR
				offsetC := ((vol.Collumns() - 1) / 2) - cC

				cordR := ir + ((vol.Rows() - r) / 2) - offsetR
				cordC := ic + ((vol.Collumns() - c) / 2) - offsetC

				if cordR < 0 || cordR > vol.Rows()-1 || cordC < 0 || cordC > vol.Collumns()-1 {
					val = 0.0
				} else {
					val = vol.GetAt(cordR, cordC, id)
				}

				sub.SetAt(ir, ic, id, val)

			}
		}
	}

	return *sub
}
func (vol *rNVolume) Equals(in rNVolume) bool {
	if !vol.EqualSize(in) {
		return false
	}

	r, c, d := vol.Dims()

	for i1 := 0; i1 < r; i1++ {
		for i2 := 0; i2 < c; i2++ {
			for i3 := 0; i3 < d; i3++ {
				if vol.GetAt(i1, i2, i3) != in.GetAt(i1, i2, i3) {
					return false
				}
			}
		}
	}

	return true
}

func (vol *rNVolume) GetAt(r, c, d int) float64 {
	return vol.Fields[d].At(r, c)
}

func (vol *rNVolume) SetAt(r, c, d int, val float64) {
	vol.Fields[d].Set(r, c, val)
}

func (vol *rNVolume) Print() {

	for i := range vol.Fields {
		fa := mat64.Formatted(&vol.Fields[i], mat64.Prefix(" "))
		fmt.Printf("Layer %v:\n\n %v\n\n", i, fa)
	}
}

// Rows of the Volume
func (vol *rNVolume) Rows() int {
	r, _, _ := vol.Dims()
	return r
}

// Collumns of the Volume
func (vol *rNVolume) Collumns() int {
	_, c, _ := vol.Dims()
	return c
}

//Depth of the Volume
func (vol *rNVolume) Depth() int {
	_, _, d := vol.Dims()
	return d
}

//EqualSize checks if the size of two volumes are the same
func (vol *rNVolume) EqualSize(a rNVolume) bool {
	i1, i2, i3 := vol.Dims()
	e1, e2, e3 := a.Dims()
	return Equal3Dim(i1, i2, i3, e1, e2, e3)
}

func (vol *rNVolume) PointReflect() {
	
}

func (vol  *rNVolume) Reflect() {
	//Kernel output

	//r,c,d := input.Dims()
	
	// for (int r= 0; r<input.Dims(); r++){
	// 	output

	// }

}
