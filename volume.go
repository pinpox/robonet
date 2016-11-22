package robonet

import (
	"errors"
	"fmt"
	"github.com/gonum/matrix/mat64"
	"log"
	"math"
	"math/rand"
)

// Volume is a basic type to hold the layer's info
type Volume struct {
	Fields []mat64.Dense
}

//SetAll sets all values of the volume from another equal-sized volume
func (vol *Volume) SetAll(v Volume) {

	if !EqualVolDim(*vol, v) {
		log.Fatal(errors.New("Volumedimensions do not match"))
	}

	*vol = v
}

//Dims returns the Dimensions of a Volume
func (vol Volume) Dims() (int, int, int) {
	d := len(vol.Fields)
	if d != 0 {
		r, c := vol.Fields[0].Dims()
		return r, c, d
	}
	return 0, 0, 0
}

//Apply applys the given kernel to the whole volume, returnung a Volume with 1 depth
func (vol *Volume) Apply(kern Kernel, strideR, strideC int) {

	r, c, _ := vol.Dims()
	r2, c2, _ := kern.Dims()

	if r%strideR != 0 || c%strideC != 0 {
		log.Fatal(errors.New("strides not applicable for this volume size"))
	}

	res := NewVolume(r/strideR, c/strideC, 1)

	for i := 0; i < r/strideR; i++ {
		for j := 0; j < c/strideC; j++ {
			res.SetAt(i, j, 0, kern.Apply(vol.SubVolumePadded(i*strideR, j*strideC, r2, c2)))
		}
	}
	//TODO normalize

	*vol = res
}

//Norm normalizes the volume to a given maximum and 0
func (vol *Volume) Norm(max float64) {

	volmin := vol.Min()

	for r := 0; r < vol.Rows(); r++ {
		for c := 0; c < vol.Collumns(); c++ {
			for d := 0; d < vol.Depth(); d++ {
				vol.SetAt(r, c, d, (vol.GetAt(r, c, d))-volmin)
			}
		}
	}
	volmin = vol.Min()
	volmax := vol.Max()

	for r := 0; r < vol.Rows(); r++ {
		for c := 0; c < vol.Collumns(); c++ {
			for d := 0; d < vol.Depth(); d++ {
				val := ((vol.GetAt(r, c, d)) * max / volmax)
				vol.SetAt(r, c, d, val)
			}
		}
	}
}

//NewVolume generates a Volume of fixed size filled with zeros
func NewVolume(r, c, d int) Volume {
	v := new(Volume)
	v.Fields = []mat64.Dense{}

	for i := 0; i < d; i++ {
		v.Fields = append(v.Fields, *mat64.NewDense(r, c, nil))
	}

	return *v
}

//NewVolumeRandom generates a Volume of fixed size filled with values between 0 and 1
func NewVolumeRandom(r, c, d int) Volume {
	v := new(Volume)
	v.Fields = []mat64.Dense{}

	for j := 0; j < d; j++ {

		data := make([]float64, r*c)
		for i := range data {
			data[i] = rand.Float64()
		}
		a := mat64.NewDense(r, c, data)

		v.Fields = append(v.Fields, *a)
	}
	return *v
}

//NewVolumeFilled generates a Volume of fixed size filled with values between 0 and 1
func NewVolumeFilled(r, c, d int, fil float64) Volume {
	v := new(Volume)
	v.Fields = []mat64.Dense{}

	for j := 0; j < d; j++ {

		data := make([]float64, r*c)
		for i := range data {
			data[i] = fil
		}
		a := mat64.NewDense(r, c, data)

		v.Fields = append(v.Fields, *a)
	}
	return *v
}

//SubVolumePadded returns a part of the original Volume. cR and cC determine the center of copying, r and c the size of the subvolume.
//If the size exceeds the underlying volume the submodule is filled(padded with Zeros.
func (vol *Volume) SubVolumePadded(cR, cC, r, c int) Volume {

	if r%2 == 0 || c%2 == 0 {
		log.Fatal(errors.New("Even dimensions not allowed for subvolumes"))
	}

	sub := NewVolume(r, c, vol.Depth())

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

	return sub
}

//SubVolume returns a part of the original Volume. tR and tC determine the center of copying, r and c the size of the subvolume.
//If the size exceeds the underlying volume the an error will be thrown, padding is not allowed.
func (vol *Volume) SubVolume(tR, tC, r, c int) Volume {

	if tR+r > vol.Rows() || tC+c > vol.Collumns() {
		log.Fatal(errors.New("Volume: Subvolume size exceeds volume dimensions"))
	}

	sub := NewVolume(r, c, vol.Depth())

	for ir := 0; ir < sub.Rows(); ir++ {
		for ic := 0; ic < sub.Collumns(); ic++ {
			for id := 0; id < sub.Depth(); id++ {
				sub.SetAt(ir, ic, id, vol.GetAt(tR+ir, tC+ic, id))
			}
		}
	}

	return sub
}

//Equals compares the volume to another volume
func (vol *Volume) Equals(in Volume) bool {
	return vol.SimilarTo(in, 0)
}

//GetAt returns the element of the volume at a given position
func (vol *Volume) GetAt(r, c, d int) float64 {
	return vol.Fields[d].At(r, c)
}

//SetAt sets the element of a volume at a given position
func (vol *Volume) SetAt(r, c, d int, val float64) {
	if r >= vol.Rows() || c >= vol.Collumns() || d >= vol.Depth() {
		//fmt.Printf("SetAt request out of bounds (RxCxD) = %vx%vx%v requested for (RxCxD) = %vx%vx%vx", r, c, d, vol.Rows(), vol.Collumns(), vol.Depth())
		panic("out od bounds")
		log.Fatal(errors.New("robonet.Volume: setAt out of bounds"))
	}
	vol.Fields[d].Set(r, c, val)
}

//Print prints the Volume to the console in a pretty format
func (vol *Volume) Print() {

	for i := range vol.Fields {
		fa := mat64.Formatted(&vol.Fields[i], mat64.Prefix(" "))
		fmt.Printf("Layer %v:\n\n %v\n\n", i, fa)
	}
}

// Rows of the Volume
func (vol *Volume) Rows() int {
	r, _, _ := vol.Dims()
	return r
}

// Collumns of the Volume
func (vol *Volume) Collumns() int {
	_, c, _ := vol.Dims()
	return c
}

//Depth of the Volume
func (vol *Volume) Depth() int {
	_, _, d := vol.Dims()
	return d
}

//EqualSize checks if the size of two volumes are the same
func (vol *Volume) EqualSize(a Volume) bool {
	i1, i2, i3 := vol.Dims()
	e1, e2, e3 := a.Dims()
	return Equal3Dim(i1, i2, i3, e1, e2, e3)
}

//PointReflect calculates the pointreflection of a volume
func (vol *Volume) PointReflect() {
	r, c, d := vol.Dims()
	temp := NewVolume(c, r, d)

	for id := 0; id < d; id++ {
		for ir := 0; ir < r; ir++ {
			for ic := 0; ic < c; ic++ {
				temp.SetAt(ic, ir, id, vol.GetAt(ir, ic, id))
			}
		}
	}
	*vol = temp
}

//Reflect calculates the reflectio of a volume (left-right)
func (vol *Volume) Reflect() {

	r, c, d := vol.Dims()
	temp := NewVolume(r, c, d)

	for id := 0; id < d; id++ {
		for ir := 0; ir < r; ir++ {
			for ic := 0; ic < c; ic++ {
				temp.SetAt(ir, ic, id, vol.GetAt(ir, c-(ic+1), id))
			}
		}
	}
	*vol = temp
}

//MulElem2 multiplies the volume with another volume element-wise
func (vol *Volume) MulElem2(v1 Volume) {
	r, c, d := vol.Dims()

	res := NewVolume(r, c, d)

	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			for k := 0; k < d; k++ {
				res.SetAt(i, j, k, vol.GetAt(i, j, k)*v1.GetAt(i, j, k))

			}
		}
	}

	*vol = res

}

//Max returns the hightest number in a volume
func (vol Volume) Max() float64 {
	max := 0.0
	for i := 0; i < vol.Rows(); i++ {
		for j := 0; j < vol.Collumns(); j++ {
			for k := 0; k < vol.Depth(); k++ {
				if vol.GetAt(i, j, k) > max {
					max = vol.GetAt(i, j, k)
				}
			}
		}
	}
	return max
}

//Min returns the lowest number in a volume
func (vol Volume) Min() float64 {
	min := 0.0
	for i := 0; i < vol.Rows(); i++ {
		for j := 0; j < vol.Collumns(); j++ {
			for k := 0; k < vol.Depth(); k++ {
				if vol.GetAt(i, j, k) < min {
					min = vol.GetAt(i, j, k)
				}
			}
		}
	}
	return min
}

//SimilarTo compares two volumes with a given threshold
func (vol Volume) SimilarTo(in Volume, threshold float64) bool {

	if !vol.EqualSize(in) {
		return false
	}

	r, c, d := vol.Dims()

	for i1 := 0; i1 < r; i1++ {
		for i2 := 0; i2 < c; i2++ {
			for i3 := 0; i3 < d; i3++ {
				if math.Abs(math.Abs(vol.GetAt(i1, i2, i3))-math.Abs(in.GetAt(i1, i2, i3))) > threshold {
					return false
				}
			}
		}
	}

	return true
}

//Elems returns the number of elements in a volume
func (vol Volume) Elems() int {
	return vol.Rows() * vol.Collumns() * vol.Depth()

}
