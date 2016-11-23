package robonet

import (
	"errors"
	"log"
)

//PoolLayer will perform a downsampling operation along the spatial dimensions (width, height), resulting in volume such as [16x16x12].
type PoolLayer struct {
	LayerFields
	SizeR   int
	SizeC   int
	StrideR int
	StrideC int
}

//Calculate for Pooling layers applies the pooling operation after the parameters have been set.
func (lay *PoolLayer) Calculate() {

	if lay.SizeR == 0 || lay.SizeC == 0 || lay.StrideR == 0 || lay.StrideC == 0 {
		log.Fatal(errors.New("PoolLayer: Parameters not set"))
	}

	if ((lay.input.Rows()%(lay.StrideR))%lay.SizeR != 0) || ((lay.input.Collumns()%(lay.StrideR))%lay.SizeC != 0) {
		log.Fatal(errors.New("PoolLayer: Input Size not divisible by factor"))
	}

	lay.output = New((lay.input.Rows()-lay.SizeR)/lay.StrideR+1, (lay.input.Collumns()-lay.SizeC)/lay.StrideC+1, lay.input.Depth())

	for r := 0; r < lay.output.Rows(); r++ {
		for c := 0; c < lay.output.Collumns(); c++ {

			res := maxPool(lay.input.SubVolume(r*lay.StrideR, c*lay.StrideC, lay.SizeR, lay.SizeC))

			for d := 0; d < lay.input.Depth(); d++ {

				lay.output.SetAt(r, c, d, res[d])
			}
		}
	}
}

//maxPool calculates the maximums of a volume's depths and returns them as slice.
//The position in the slice matches the depth of that value in the input volume
func maxPool(vol Volume) (res []float64) {

	for d := 0; d < vol.Depth(); d++ {

		max := 0.0

		for r := 0; r < vol.Rows(); r++ {
			for c := 0; c < vol.Collumns(); c++ {
				if tmp := vol.GetAt(r, c, d); tmp > max {
					max = tmp
				}
			}
		}
		res = append(res, max)
	}
	return res
}
