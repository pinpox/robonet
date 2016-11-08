package main

import (
	"math"
)

//SigmoidFast calcultes the value for activation using a fast sigmoid approximation
func SigmoidFast(x float64) float64 {
	return x / (1 + math.Abs(x))
}

//Equal3Dim checks if the size of two volumes are the same
func Equal3Dim(e1, e2, e3, i1, i2, i3 int) bool {
	return (e1 == i1 && e2 == i2 && e3 == i3)
}
