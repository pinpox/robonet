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

func (vol rNVolume) Dims() (int, int, int) {
	i1, i2 := vol.Fields[0].Dims()
	i3 := len(vol.Fields)
	return i1, i2, i3
}

//Apply applys the given filter to the whole volume, returnung a Volume with 1 depth
func (vol rNVolume) Apply(f Filter) rNVolume {

	//Check correct output
	_, _, a := vol.Dims()
	if a != 1 {
		panic("should have returned a plane (1dim)")
	}

	//TODO apply the filter to the volume
	return vol
}

//NewRNVolume generates a rNVolume of fixed size filled with zeros
func NewRNVolume(h int, w int, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *mat64.NewDense(h, w, nil))
	}
	return v
}

//NewRNVolumeRandom generates a rNVolume of fixed size filled with values between 0 and 1
func NewRNVolumeRandom(h int, w int, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	data := make([]float64, w*h)
	for i := range data {
		data[i] = rand.Float64()
	}
	a := mat64.NewDense(w, h, data)

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *a)
	}
	return v
}

//SubVolumePadded returns a part of the original Volume. i and j determine the center of copying, width and height the size of the subvolume.
//If the size exceeds the underlying volume the submodule is filled(padded with Zeros.
func (vol rNVolume) SubVolumePadded(i, j, height, width int) rNVolume {
	//TODO
	return *new(rNVolume)
}

func (vol rNVolume) Height() int {
	return len(vol.Fields)
}

func (vol rNVolume) Print() {

	for i := range vol.Fields {
		fa := mat64.Formatted(&vol.Fields[i], mat64.Prefix(" "))
		fmt.Printf("Layer %v:\n\n %v\n\n", i, fa)
	}
}

func (vol *rNVolume) Width() int {
	_, c := vol.Fields[0].Dims()
	return c
}

func (vol rNVolume) Depth() int {
	r, _ := vol.Fields[0].Dims()
	return r
}

//Equal3Dim checks if the size of two volumes are the same
func (vol rNVolume) EqualSize(a rNVolume) bool {
	i1, i2, i3 := vol.Dims()
	e1, e2, e3 := a.Dims()
	return Equal3Dim(i1, i2, i3, e1, e2, e3)
}
