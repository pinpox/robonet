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

func EqualNDim(v1, v2 []int) bool {
	if len(v1) == len(v2) {
		for k, v := range v1 {
			if v != v2[k] {
				return false
			}
		}
		return true
	}
	return false
}

//Odd3Dim checks if the rows and collumns are odd
func Odd3Dim(i1, i2, i3 int) bool {
	return !(i1%2 == 0 && i2%2 == 0)
}

//EqualVolDim checks if two given volumes have the same dimensions
func EqualVolDim(v1, v2 Volume) bool {
	if len(v1.Shape()) == len(v2.Shape()) {
		for k, v := range v1.Shape() {
			if v != v2.Shape()[k] {
				return false
			}
		}
		return true
	}
	return false
}

//VolumeFromTIFF creates a volume from a given file
func VolumeFromTIFF(path string) Volume {

	file, err := os.Open(path)
	if err != nil {
		panic("could not read")
	}
	// decode jpeg into image.Image
	img, err := tiff.Decode(file)
	if err != nil {
		panic("could not decode")
	}
	defer file.Close()
	return ImageToVolume(img)
}

//VolumeFromJPEG creates a volume from a given file
func VolumeFromJPEG(path string) Volume {

	file, err := os.Open(path)
	if err != nil {
		panic("could not read")
	}
	img, err := jpeg.Decode(file)
	if err != nil {
		panic("could not decode")
	}
	defer file.Close()
	return ImageToVolume(img)
}

//SaveVolumeToTIFF saves a volume to a given TIFF-file
func SaveVolumeToTIFF(path string, vol Volume) {

	toimg, _ := os.Create(path)
	defer toimg.Close()
	m := VolumeToImage(vol)
	tiff.Encode(toimg, m, nil)
}

//SaveVolumeToJPEG saves a volume to a given JPEG-file
func SaveVolumeToJPEG(path string, vol Volume) {

	toimg, _ := os.Create(path)
	defer toimg.Close()
	m := VolumeToImage(vol)
	jpeg.Encode(toimg, m, nil)
}

//ImageToVolume creates a volume from a image.Image
func ImageToVolume(img image.Image) Volume {

	vol := New(img.Bounds().Max.X, img.Bounds().Max.Y, 3)

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

//VolumeToImage converts a volume to a image. Values are rounded to 2 decimal palaces
func VolumeToImage(vol Volume) image.Image {

	m := image.NewRGBA(image.Rect(0, 0, vol.Rows(), vol.Collumns()))
	switch vol.Depth() {
	case 1:

		if Round(vol.Max(), 2) > 765 {
			log.Fatal("Can't save volume as BW image, has values over 765: ", vol.Max())

		}

		for c := 0; c < vol.Collumns(); c++ {
			for r := 0; r < vol.Rows(); r++ {

				red, green, blue, alpha := uint8(Round(vol.GetAt(r, c, 0), 2)), uint8(Round(vol.GetAt(r, c, 0), 2)), uint8(Round(vol.GetAt(r, c, 0), 2)), uint8(255)
				//fmt.Println("setting ", red, blue, green)
				m.Set(r, c, color.RGBA{red, green, blue, alpha})
			}
		}
	case 3:

		if Round(vol.Max(), 2) > 255 {
			log.Fatal("Can't save volume as color image, has values over 255: ", vol.Max())

		}
		for c := 0; c < vol.Collumns(); c++ {
			for r := 0; r < vol.Rows(); r++ {

				red, green, blue, alpha := uint8(Round(vol.GetAt(r, c, 0), 2)), uint8(Round(vol.GetAt(r, c, 1), 2)), uint8(Round(vol.GetAt(r, c, 2), 2)), uint8(255)
				m.Set(r, c, color.RGBA{red, green, blue, alpha})
			}
		}
	default:
		log.Fatal(errors.New("only 3-deep or 1-deep volumes can be saved to images"))
	}
	return m
}

//CompareJPEG compares two  JPEGs pixel-wise. A threshold (0-255) is specified. 0 means the two images are identical
func CompareJPEG(path1, path2 string, threshold float64) bool {
	imgvol1 := VolumeFromJPEG(path1)
	imgvol2 := VolumeFromJPEG(path2)
	return imgvol1.SimilarTo(imgvol2, threshold)
}

//CompareTIFF compares two  TIFFs pixel-wise. A threshold (0-255) is specified. 0 means the two images are identical
func CompareTIFF(path1, path2 string, threshold float64) bool {
	imgvol1 := VolumeFromTIFF(path1)
	imgvol2 := VolumeFromTIFF(path2)
	return imgvol1.SimilarTo(imgvol2, threshold)
}

//Round rounds to a given number of places
func Round(val float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= .5 {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}
