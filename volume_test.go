package robonet

import "fmt"
import "testing"
import "github.com/gonum/matrix/mat64"

var volumeSizes = [][]int{
	{1, 1, 1},
	{3, 3, 3},
	{5, 5, 5},
	{5, 5, 1},
	{5, 1, 5},
	{1, 1, 5},
	{5, 1, 5},
	{1, 2, 3},
	{3, 2, 1},
}

var data1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
var data2 = []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
var data3 = []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
var testVol = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, data1), *mat64.NewDense(3, 3, data2), *mat64.NewDense(3, 3, data3)}}

// Check if Correct Dimensions are displayed
func TestDims(t *testing.T) {
	for _, v := range volumeSizes {
		vol := new(Volume)

		for i := 0; i < v[2]; i++ {
			vol.Fields = append(vol.Fields, *mat64.NewDense(v[0], v[1], nil))
		}

		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}

}

// Check correctly sized volume is created
func TestNewRNVolume(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewRNVolume(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}

// Check correctly sized volume is created
func TestNewRNVolumeRandom(t *testing.T) {
	for _, v := range volumeSizes {
		vol := NewRNVolumeRandom(v[0], v[1], v[2])
		i1, i2, i3 := vol.Dims()
		if !Equal3Dim(i1, i2, i3, v[0], v[1], v[2]) {
			t.Error("Expected ", v[0], v[1], v[2], ", got ", i1, i2, i3)
		}
	}
}

// Check if Subvolumes are correctly extractes and padded
func TestSubVolumePadded(t *testing.T) {

	var sub1 = []float64{
		0, 0, 0,
		0, 0, 1,
		0, 3, 4}

	var sub2 = []float64{
		0, 0, 0,
		0, 9, 10,
		0, 12, 13}

	var sub3 = []float64{
		0, 0, 0,
		0, 18, 19,
		0, 21, 22}

	var subVol1 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}
	sub1 = []float64{
		0, 0, 0,
		0, 1, 2,
		3, 4, 5}

	sub2 = []float64{
		0, 0, 0,
		9, 10, 11,
		12, 13, 14}

	sub3 = []float64{
		0, 0, 0,
		18, 19, 20,
		21, 22, 23}

	var subVol2 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		0, 0, 0,
		1, 2, 0,
		4, 5, 0}

	sub2 = []float64{
		0, 0, 0,
		10, 11, 0,
		13, 14, 0}

	sub3 = []float64{
		0, 0, 0,
		19, 20, 0,
		22, 23, 0}

	var subVol3 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		0, 0, 1,
		0, 3, 4,
		0, 6, 7}

	sub2 = []float64{
		0, 9, 10,
		0, 12, 13,
		0, 15, 16}

	sub3 = []float64{
		0, 18, 19,
		0, 21, 22,
		0, 24, 25}

	var subVol4 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = data1
	sub2 = data2
	sub3 = data3

	var subVol5 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		4, 5, 0,
		7, 8, 0,
		0, 0, 0}

	sub2 = []float64{
		13, 14, 0,
		16, 17, 0,
		0, 0, 0}

	sub3 = []float64{
		22, 23, 0,
		25, 26, 0,
		0, 0, 0}

	var subVol6 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{0}
	sub2 = []float64{9}
	sub3 = []float64{18}
	var subVol7 = Volume{Fields: []mat64.Dense{*mat64.NewDense(1, 1, sub1), *mat64.NewDense(1, 1, sub2), *mat64.NewDense(1, 1, sub3)}}

	sub1 = []float64{
		0, 0, 0, 0, 0,
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0,
		0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 0, 0, 0, 0,
		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0,
		0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 0, 0, 0, 0,
		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0,
		0, 0, 0, 0, 0}
	var subVol8 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 5, sub1), *mat64.NewDense(5, 5, sub2), *mat64.NewDense(5, 5, sub3)}}

	sub1 = []float64{
		0, 0, 0,
		0, 1, 2,
		3, 4, 5,
		6, 7, 8,
		0, 0, 0}

	sub2 = []float64{
		0, 0, 0,
		9, 10, 11,
		12, 13, 14,
		15, 16, 17,
		0, 0, 0}

	sub3 = []float64{
		0, 0, 0,
		18, 19, 20,
		21, 22, 23,
		24, 25, 26,
		0, 0, 0}

	var subVol9 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 3, sub1), *mat64.NewDense(5, 3, sub2), *mat64.NewDense(5, 3, sub3)}}

	sub1 = []float64{
		0, 0, 1, 2, 0,
		0, 3, 4, 5, 0,
		0, 6, 7, 8, 0}

	sub2 = []float64{
		0, 9, 10, 11, 0,
		0, 12, 13, 14, 0,
		0, 15, 16, 17, 0}

	sub3 = []float64{
		0, 18, 19, 20, 0,
		0, 21, 22, 23, 0,
		0, 24, 25, 26, 0}

	var subVol10 = Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, sub1), *mat64.NewDense(3, 5, sub2), *mat64.NewDense(3, 5, sub3)}}

	sub1 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 1, 2, 0, 0,
		0, 0, 3, 4, 5, 0, 0,
		0, 0, 6, 7, 8, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 9, 10, 11, 0, 0,
		0, 0, 12, 13, 14, 0, 0,
		0, 0, 15, 16, 17, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 18, 19, 20, 0, 0,
		0, 0, 21, 22, 23, 0, 0,
		0, 0, 24, 25, 26, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol11 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 7, sub1), *mat64.NewDense(5, 7, sub2), *mat64.NewDense(5, 7, sub3)}}

	sub1 = []float64{
		0, 0, 1, 2, 0, 0, 0,
		0, 3, 4, 5, 0, 0, 0,
		0, 6, 7, 8, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub2 = []float64{
		0, 9, 10, 11, 0, 0, 0,
		0, 12, 13, 14, 0, 0, 0,
		0, 15, 16, 17, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	sub3 = []float64{
		0, 18, 19, 20, 0, 0, 0,
		0, 21, 22, 23, 0, 0, 0,
		0, 24, 25, 26, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0}

	var subVol12 = Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 7, sub1), *mat64.NewDense(5, 7, sub2), *mat64.NewDense(5, 7, sub3)}}

	res1 := testVol.SubVolumePadded(0, 0, 3, 3)
	res2 := testVol.SubVolumePadded(0, 1, 3, 3)
	res3 := testVol.SubVolumePadded(0, 2, 3, 3)
	res4 := testVol.SubVolumePadded(1, 0, 3, 3)
	res5 := testVol.SubVolumePadded(1, 1, 3, 3)
	res6 := testVol.SubVolumePadded(2, 2, 3, 3)
	res7 := testVol.SubVolumePadded(0, 0, 1, 1)
	res8 := testVol.SubVolumePadded(1, 1, 5, 5)
	res9 := testVol.SubVolumePadded(1, 1, 5, 3)
	res10 := testVol.SubVolumePadded(1, 1, 3, 5)
	res11 := testVol.SubVolumePadded(1, 1, 5, 7)
	res12 := testVol.SubVolumePadded(2, 2, 5, 7)

	// 0, 1, 2
	// 3, 4, 5
	// 6, 7, 8

	// 9, 10, 11
	// 12, 13,14
	// 15,16, 17

	//18, 19, 20
	//21, 22, 23
	//24, 25, 26

	if !res1.Equals(subVol1) {
		t.Error("Expected", subVol1, ", got ", res1)
		fmt.Println("Expected 1")
		subVol1.Print()
		fmt.Println("got")
		res1.Print()
	}
	if !res2.Equals(subVol2) {
		t.Error("Expected", subVol2, ", got ", res2)
		fmt.Println("Expected 2")
		subVol2.Print()
		fmt.Println("got")
		res2.Print()
	}

	if !res3.Equals(subVol3) {
		t.Error("Expected", subVol3, ", got ", res3)
		fmt.Println("Expected 3")
		subVol3.Print()
		fmt.Println("got")
		res3.Print()
	}

	if !res4.Equals(subVol4) {
		t.Error("Expected", subVol4, ", got ", res4)
		fmt.Println("Expected 4")
		subVol4.Print()
		fmt.Println("got")
		res4.Print()
	}

	if !res5.Equals(subVol5) {
		t.Error("Expected", subVol5, ", got ", res5)
		fmt.Println("Expected 5")
		subVol5.Print()
		fmt.Println("got")
		res5.Print()
	}

	if !res6.Equals(subVol6) {
		t.Error("Expected", subVol6, ", got ", res6)
		fmt.Println("Expected 6")
		subVol6.Print()
		fmt.Println("got")
		res6.Print()
	}

	if !res7.Equals(subVol7) {
		t.Error("Expected", subVol7, ", got ", res7)
		fmt.Println("Expected 7")
		subVol7.Print()
		fmt.Println("got")
		res7.Print()
	}

	if !res8.Equals(subVol8) {
		t.Error("Expected", subVol8, ", got ", res8)
		fmt.Println("Expected 8")
		subVol8.Print()
		fmt.Println("got")
		res8.Print()
	}

	if !res9.Equals(subVol9) {
		t.Error("Expected", subVol9, ", got ", res9)
		fmt.Println("Expected 9")
		subVol9.Print()
		fmt.Println("got")
		res9.Print()
	}

	if !res10.Equals(subVol10) {
		t.Error("Expected", subVol10, ", got ", res10)
		fmt.Println("Expected 10")
		subVol10.Print()
		fmt.Println("got")
		res10.Print()
	}

	if !res11.Equals(subVol11) {
		t.Error("Expected", subVol11, ", got ", res11)
		fmt.Println("Expected 11")
		subVol11.Print()
		fmt.Println("got")
		res11.Print()
	}

	if !res12.Equals(subVol12) {
		t.Error("Expected", subVol12, ", got ", res12)
		fmt.Println("Expected 12")
		subVol12.Print()
		fmt.Println("got")
		res12.Print()
	}
	/*}
	  func TestVolumeApply(t *testing.T) {

	  	//Create new vol
	  	//Creeat kernl
	  	//apply kernl
	  	//test dims
	  	//test nums

	  	ker := NewKernel(3, 3, 3)
	  	res := testVol.Apply(ker)

	  	resExp := nil //TODO

	  	if !res.Equals(resExp) {
	  		t.Error("Result incorrect")
	  		fmt.Println("Expected")
	  		resExp.Print()
	  		fmt.Println("Result")
	  		res.Print()
	  	}
	*/
}

func TestVolumeReflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1 := NewKernel(3, 3, 3)
	kern1.SetAll(testVol1)

	layer1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	layer2 = []float64{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	testVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}

	kern2 := NewKernel(3, 5, 2)
	kern2.SetAll(testVol2)

	//Result
	layer1 = []float64{2, 1, 0, 5, 4, 3, 8, 7, 6}
	layer2 = []float64{11, 10, 9, 14, 13, 12, 17, 16, 15}
	layer3 = []float64{20, 19, 18, 23, 22, 21, 26, 25, 24}
	testVol1Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{4, 3, 2, 1, 0, 9, 8, 7, 6, 5, 14, 13, 12, 11, 10}
	layer2 = []float64{19, 18, 17, 16, 15, 24, 23, 22, 21, 20, 29, 28, 27, 26, 25}
	testVol2Reflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}
	//Compare
	testVol1.Reflect()
	if !(testVol1Reflected.Equals(testVol1)) {

		t.Error("Reflect () Expected", testVol1Reflected, " got", testVol1)
		testVol1Reflected.Print()
		testVol1.Print()
	}

	testVol2.Reflect()
	if !(testVol2Reflected.Equals(testVol2)) {

		t.Error("Reflect () Expected", testVol2Reflected, " got", testVol2)
		testVol2Reflected.Print()
		testVol2.Print()
	}
}

func TestVolumePointReflect(t *testing.T) {

	//Input
	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol1 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	kern1 := NewKernel(3, 3, 3)
	kern1.SetAll(testVol1)

	layer1 = []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}
	layer2 = []float64{15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29}
	testVol2 := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, layer1), *mat64.NewDense(3, 5, layer2)}}

	kern2 := NewKernel(3, 5, 2)
	kern2.SetAll(testVol2)

	//Result
	layer1 = []float64{0, 3, 6, 1, 4, 7, 2, 5, 8}
	layer2 = []float64{9, 12, 15, 10, 13, 16, 11, 14, 17}
	layer3 = []float64{18, 21, 24, 19, 22, 25, 20, 23, 26}
	testVol1PointReflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	layer1 = []float64{0, 5, 10, 1, 6, 11, 2, 7, 12, 3, 8, 13, 4, 9, 14}
	layer2 = []float64{15, 20, 25, 16, 21, 26, 17, 22, 27, 18, 23, 28, 19, 24, 29}
	testVol2PointReflected := Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 3, layer1), *mat64.NewDense(5, 3, layer2)}}
	//Compare
	testVol1.PointReflect()
	if !(testVol1PointReflected.Equals(testVol1)) {

		t.Error("Reflect () Expected", testVol1PointReflected, " got", testVol1)
		testVol1PointReflected.Print()
		testVol1.Print()
	}

	testVol2.PointReflect()
	if !(testVol2PointReflected.Equals(testVol2)) {

		t.Error("Reflect () Expected", testVol2PointReflected, " got", testVol2)
		testVol2PointReflected.Print()
		testVol2.Print()
	}

}

func TestVolumeApply(t *testing.T) {

	//TODO

	//    layer1 = []float64{0,3,6,1,4,7,2,5,8}
	// layer2 = []float64{9,12,15,10,13,16,11,14,17}
	// layer3 = []float64{18,21,24,19,22,25,20,23,26}
	// testVol := Volume{Fields: []mat64.Dense{*mat64.NewDense(5, 5, layer1), *mat64.NewDense(5, 5, layer2), *mat64.NewDense(5, 5, layer3)}}

}

func TestVolumeMax(t *testing.T) {

	layer1 := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8}
	layer2 := []float64{9, 10, 11, 12, 13, 14, 15, 16, 17}
	layer3 := []float64{18, 19, 20, 21, 22, 23, 24, 25, 26}
	testVol := Volume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, layer1), *mat64.NewDense(3, 3, layer2), *mat64.NewDense(3, 3, layer3)}}

	if !(testVol.Max() == 26) {
		t.Error("expected 29 got ", testVol.Max())
	}
}
