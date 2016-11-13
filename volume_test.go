package robonet

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
var testVol = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, data1), *mat64.NewDense(3, 3, data2), *mat64.NewDense(3, 3, data3)}}

// Check if Correct Dimensions are displayed
func TestDims(t *testing.T) {
	for _, v := range volumeSizes {
		vol := new(rNVolume)

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

	var subVol1 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

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

	var subVol2 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

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

	var subVol3 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	// 0, 1, 2
	// 3, 4, 5
	// 6, 7, 8

	// 9, 10, 11
	// 12, 13,14
	// 15,16, 17

	//18, 19, 20
	//21, 22, 23
	//24, 25, 26

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

	var subVol4 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = data1
	sub2 = data2
	sub3 = data3

	var subVol5 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{
		4, 5, 0,
		7, 8, 0,
		0, 0, 0}

	sub2 = []float64{
		13, 14, 0,
		16, 17, 0,
		0, 0, 0}

	sub3 = []float64{
		21, 22, 0,
		24, 25, 0,
		0, 0, 0}

	var subVol6 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 3, sub1), *mat64.NewDense(3, 3, sub2), *mat64.NewDense(3, 3, sub3)}}

	sub1 = []float64{0}
	sub2 = []float64{9}
	sub3 = []float64{18}
	var subVol7 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(1, 1, sub1), *mat64.NewDense(1, 1, sub2), *mat64.NewDense(1, 1, sub3)}}

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
	var subVol8 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(5, 5, sub1), *mat64.NewDense(5, 5, sub2), *mat64.NewDense(5, 5, sub3)}}

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

	var subVol9 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(5, 3, sub1), *mat64.NewDense(5, 3, sub2), *mat64.NewDense(5, 3, sub3)}}

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

	var subVol10 = rNVolume{Fields: []mat64.Dense{*mat64.NewDense(3, 5, sub1), *mat64.NewDense(3, 5, sub2), *mat64.NewDense(3, 5, sub3)}}

	res1 := testVol.SubVolumePadded(0, 0, 3, 3)
	res2 := testVol.SubVolumePadded(1, 0, 3, 3)
	res3 := testVol.SubVolumePadded(2, 0, 3, 3)
	res4 := testVol.SubVolumePadded(0, 1, 3, 3)
	res5 := testVol.SubVolumePadded(1, 1, 3, 3)
	res6 := testVol.SubVolumePadded(2, 2, 3, 3)
	res7 := testVol.SubVolumePadded(0, 0, 1, 1)
	res8 := testVol.SubVolumePadded(1, 1, 5, 5)
	res9 := testVol.SubVolumePadded(2, 2, 5, 3)
	res10 := testVol.SubVolumePadded(2, 2, 3, 5)

	if !res1.Equals(subVol1) {
		t.Error("Expected", subVol1, ", got ", res1)
		//subVol1.Print()
		//res1.Print()
	}

	if !res2.Equals(subVol2) {
		t.Error("Expected", subVol2, ", got ", res2)
		//subVol2.Print()
		//res2.Print()
	}

	if !res3.Equals(subVol3) {
		t.Error("Expected", subVol3, ", got ", res3)
		//subVol3.Print()
		//res3.Print()
	}

	if !res4.Equals(subVol4) {
		t.Error("Expected", subVol4, ", got ", res4)
		//subVol4.Print()
		//res4.Print()
	}

	if !res5.Equals(subVol5) {
		t.Error("Expected", subVol5, ", got ", res5)
		//subVol5.Print()
		//res5.Print()
	}

	if !res6.Equals(subVol6) {
		t.Error("Expected", subVol6, ", got ", res6)
		//subVol6.Print()
		//res6.Print()
	}

	if !res7.Equals(subVol7) {
		t.Error("Expected", subVol7, ", got ", res7)
		//subVol7.Print()
		//res7.Print()
	}

	if !res8.Equals(subVol8) {
		t.Error("Expected", subVol8, ", got ", res8)
		//subVol8.Print()
		//res8.Print()
	}

	if !res9.Equals(subVol9) {
		t.Error("Expected", subVol9, ", got ", res9)
		//subVol9.Print()
		//res9.Print()
	}

	if !res10.Equals(subVol10) {
		t.Error("Expected", subVol10, ", got ", res10)
		//subVol10.Print()
		//res10.Print()
	}
}
