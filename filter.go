package robonet

import (
	"fmt"
)

// Filter represets a basic conv filter
type Filter struct {
	values rNVolume
}

//NewFilter creates a new filter initialized with zeros
func NewFilter(height int, width, depth int) *Filter {
	if !Odd3Dim(height, width, depth) {
		panic("Filter must have odd width and heigth")
	}
	g := Filter{*NewRNVolume(height, width, depth)}
	return &g
}

//NewFilterRandom creates a new filter initialized with random values
func NewFilterRandom(height int, width, depth int) *Filter {
	if Odd3Dim(height, width, depth) {
		panic("Filter must have odd width and heigth")
	}
	g := Filter{*NewRNVolumeRandom(height, width, depth)}
	return &g
}

//Print shows show the filter's matrix string representation
func (f Filter) Print() {
	f.values.Print()
}

//Dims returns the  size of the filter
func (f Filter) Dims() (int, int, int) {
	return f.values.Dims()
}

//Apply applys the filter to a equally sized chunk of a volume
//Only filters of the same size as the volume can be applied
func (f Filter) Apply(in rNVolume) float64 {

	float64 ConvResult

	if !(f.values.EqualSize(in)) {
		fmt.Println("Filter size doesn't match input")
		panic("Filter size doesn't match input")
	}

	// 1) reflect kernel
	// 2) multiply pairwise
	// 3) normalize

	return ConvResult
}

func (f Filter) pointReflection (input Filter) Filter{
	//TO DO
}

func (f Filter) reflection (input Filter) Filter{
	//TO DO
}
