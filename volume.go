package main

import (
	"fmt"
	"github.com/gonum/matrix/mat64"
	"math/rand"
)

// rNVolume is a basic type to hold the layer's information
type rNVolume struct {
	Fields []mat64.Dense
}

//newRNVolume generates a rNVolume of fixed size filled with zeros
func newRNVolume(h int, w int, d int) *rNVolume {
	v := new(rNVolume)
	v.Fields = []mat64.Dense{}

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *mat64.NewDense(h, w, nil))
	}
	return v
}

//newRNVolumeRandom generates a rNVolume of fixed size filled with values between 0 and 1
func newRNVolumeRandom(h int, w int, d int) *rNVolume {
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

func (v rNVolume) Height() int {
	return len(v.Fields)
}

func (v rNVolume) Print() {

	for i := range v.Fields {
		fa := mat64.Formatted(&v.Fields[i], mat64.Prefix(" "))
		fmt.Printf("Layer %v:\n\n %v\n\n", i, fa)
	}
}

func (v *rNVolume) Width() int {
	_, c := v.Fields[0].Dims()
	return c
}

func (v rNVolume) Depth() int {
	r, _ := v.Fields[0].Dims()
	return r
}
