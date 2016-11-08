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
