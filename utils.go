package robonet

import (
	"errors"
	"golang.org/x/image/tiff"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
)

//SigmoidFast calcultes the value for activation using a fast sigmoid approximation
func SigmoidFast(x float64) float64 {
	return x / (1 + math.Abs(x))
}

//Equal3Dim checks if the size of two volumes are the same
func Equal3Dim(e1, e2, e3, i1, i2, i3 int) bool {
	return (e1 == i1 && e2 == i2 && e3 == i3)
}

//Odd3Dim checks if the rows and collumns are odd
func Odd3Dim(i1, i2, i3 int) bool {
	return !(i1%2 == 0 && i2%2 == 0)
}

//EqualVolDim checks if two given volumes have the same dimensions
func EqualVolDim(v1, v2 Volume) bool {
	i1, i2, i3 := v1.Dims()
	e1, e2, e3 := v2.Dims()

	return Equal3Dim(i1, i2, i3, e1, e2, e3)
}

//VolumeFromImageFile creates a volume from a given file
func VolumeFromImageFile(path string) Volume {

	file, err := os.Open(path)
	if err != nil {
		panic("could not read")
	}
	// decode jpeg into image.Image
	img, err := jpeg.Decode(file)
	if err != nil {
		panic("could not decode")
	}
	file.Close()
	vol := NewVolume(img.Bounds().Max.X, img.Bounds().Max.Y, 3)

	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			r, g, b, _ := img.At(x, y).RGBA()
			r, g, b = (r*255)/65535, (g*255)/65535, (b*255)/65535
			vol.SetAt(x, y, 0, float64(r))
			vol.SetAt(x, y, 1, float64(g))
			vol.SetAt(x, y, 2, float64(b))
		}
	}
	return vol
}

//SaveVolumeToFile saves a volume to a given file
func SaveVolumeToFile(path string, vol Volume) {

	if vol.Depth() != 3 {
		log.Fatal(errors.New("only 3-deep volumes can be saved to images"))
	}

	toimg, _ := os.Create(path)
	defer toimg.Close()

	m := image.NewRGBA(image.Rect(0, 0, vol.Rows(), vol.Collumns()))
	for c := 0; c < vol.Collumns(); c++ {
		for r := 0; r < vol.Rows(); r++ {

			red, green, blue, alpha := uint8(vol.GetAt(r, c, 0)), uint8(vol.GetAt(r, c, 1)), uint8(vol.GetAt(r, c, 2)), uint8(255)
			m.Set(r, c, color.RGBA{red, green, blue, alpha})
		}
	}
	tiff.Encode(toimg, m, nil)
	//jpeg.Encode(toimg, m, nil)
}

//SaveVolumeToFileBW saves a volume to a given file in black and white
func SaveVolumeToFileBW(path string, vol Volume) {

	if vol.Depth() != 1 {
		log.Fatal(errors.New("only 1-deep volumes can be saved to BW-images"))
	}
	toimg, _ := os.Create(path)
	defer toimg.Close()

	m := image.NewRGBA(image.Rect(0, 0, vol.Rows(), vol.Collumns()))
	for c := 0; c < vol.Collumns(); c++ {
		for r := 0; r < vol.Rows(); r++ {

			red, green, blue, alpha := uint8(vol.GetAt(r, c, 0)), uint8(vol.GetAt(r, c, 0)), uint8(vol.GetAt(r, c, 0)), uint8(255)
			m.Set(r, c, color.RGBA{red, green, blue, alpha})
		}
	}
	tiff.Encode(toimg, m, nil)
	//jpeg.Encode(toimg, m, nil)
}
