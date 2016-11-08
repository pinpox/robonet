package main

import (
	"fmt"
	//"github.com/gonum/matrix/mat64"
)

// Filter represets a basic conv filter
type Filter struct {
	values rNVolume
}

//NewFilter creates a new filter initialized with zeros
func NewFilter(height int, width, depth int) *Filter {
	g := Filter{*newRNVolume(height, width, depth)}
	return &g
}

//NewFilterRandom creates a new filter initialized with random values
func NewFilterRandom(height int, width, depth int) *Filter {
	g := Filter{*newRNVolumeRandom(height, width, depth)}
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
func (f Filter) Apply(in rNVolume) float64 {

	if !(f.values.EqualSize(in)) {
		fmt.Println("Filter size doesn't match input")
		panic("Filter size doesn't match input")
	}

	//TODO calc
	//TODO normalize

	//a := mat64.NewDense(2, 4, []float64{
	//1, 2, 3, 4,
	//5, 6, 7, 8,
	//})
	//b := mat64.NewDense(4, 3, []float64{
	//1, 2, 3,
	//4, 5, 6,
	//7, 8, 9,
	//10, 11, 12,
	//})
	//var m mat64.Dense
	////var n mat64.Dense
	//m.Mul(a, b)
	////normFactor := n.Sum()

	//fmt.Println("multiplication output")
	//fmt.Println(mat64.Formatted(&m))

	return 1

}
