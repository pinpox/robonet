package robonet

import "testing"
import "github.com/gonum/matrix/mat64"


func TestNewKernel(t *testing.T) {
	//TODO
}

func TestNewKernelRandom(t *testing.T) {
	//TODO
}

func TestPrint(t *testing.T) {
	//TODO
}

func TestFilterDims(t *testing.T) {
	//TODO
}

func TestApply(t *testing.T) {
	//TODO
}

func TestPointReflect(t *testing.T) {
	//Input
	// layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	// layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	// layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	// testVol1 := rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	// kern1 := NewKernel(3,3,3)
	// kern1.SetAll(testVol1)

	// //Result
	// layer1 = []float64{8,7,6,5,4,3,2,1,0}
	// layer2 = []float64{17,16,15,14,13,12,11,10,9}
	// layer3 = []float64{26,25,24,23,22,21,20,19,18}
	// testVol1PointReflected := rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	// //Compare
	// testVol1.PointReflect()
	// if !(testVol1PointReflected.Equals(testVol1)) {
	// 	t.Error("Expected" , testVol1PointReflected , " got" , testVol1)
	// }
}

func TestKernelReflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1 := NewKernel(3,3,3)
	kern1.SetAll(testVol1)

	//Result
	layer1 = []float64{2, 1, 0, 5, 4, 3, 8, 7, 6}
	layer2 = []float64{11, 10, 9, 14, 13, 12, 17, 16 ,15}
	layer3 = []float64{20,19,18,23,22,21,26,25,24}
	testVol1Reflected := rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1Reflected := NewKernel(3,3,3)
	kern1Reflected.SetAll(testVol1Reflected)

	//Compare
	kern1.Reflect()
	if !(kern1Reflected.Equals(kern1)) {

		t.Error("Reflect () Expected" , kern1Reflected , " got" , kern1)
		kern1Reflected.Print()
		kern1.Print()
	}
}

