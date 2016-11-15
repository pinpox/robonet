# robonet
--
    import "gitlab.com/binaryplease/robonet"


## Usage

#### func  Equal3Dim

```go
func Equal3Dim(e1, e2, e3, i1, i2, i3 int) bool
```
Equal3Dim checks if the size of two volumes are the same

#### func  EqualVolDim

```go
func EqualVolDim(v1, v2 Volume) bool
```
EqualVolDim checks if two given volumes have the same dimensions

#### func  Odd3Dim

```go
func Odd3Dim(i1, i2, i3 int) bool
```
Odd3Dim checks if the rows and collumns are odd

#### func  SigmoidFast

```go
func SigmoidFast(x float64) float64
```
SigmoidFast calcultes the value for activation using a fast sigmoid
approximation

#### type ConvLayer

```go
type ConvLayer struct {
	Layer
}
```

ConvLayer basic type for a convolutional layer The layer will compute the output
of neurons that are connected to local regions in the input, each computing a
dot product between their weights and a small region they are connected to in
the input volume. This may result in volume such as [32x32x12] if we decided to
use 12 filters.

#### func (*ConvLayer) AddKernel

```go
func (l *ConvLayer) AddKernel(kern Kernel, strideR, strideC int)
```
AddKernel adds a kernel to a layer

#### func (*ConvLayer) Calculate

```go
func (l *ConvLayer) Calculate(vol Volume) Volume
```
Calculate applys all Kernels to a given Volume

#### func (ConvLayer) Kernels

```go
func (l ConvLayer) Kernels() []Kernel
```
Kernels returns the kernels of the layer

#### type FCLayer

```go
type FCLayer struct {
	Layer
}
```

FCLayer (i.e. fully-connected) layer will compute the class scores, resulting in
volume of size [1x1x10], where each of the 10 numbers correspond to a class
score, such as among the 10 categories of CIFAR-10. As with ordinary Neural
Networks and as the name implies, each neuron in this layer will be connected to
all the numbers in the previous volume.

#### type InputLayer

```go
type InputLayer struct {
	Layer
}
```

InputLayer [32x32x3] will hold the raw pixel values of the image, in this case
an image of width 32, height 32, and with three color channels R,G,B.

#### func (*InputLayer) SetInput

```go
func (l *InputLayer) SetInput(in Volume)
```
SetInput sets the input of a input layer

#### type Kernel

```go
type Kernel struct {
	Volume
}
```

Kernel represets a basic conv kernel

#### func  NewKernel

```go
func NewKernel(r, c, d int) Kernel
```
NewKernel creates a new kernel initialized with zeros

#### func  NewKernelRandom

```go
func NewKernelRandom(r, c, d int) *Kernel
```
NewKernelRandom creates a new kernel initialized with random values

#### func (Kernel) Apply

```go
func (kern Kernel) Apply(in Volume) float64
```
Apply applys the kernel to a equally sized chunk of a volume Only kernels of the
same size as the volume can be applied

#### func (*Kernel) Equals

```go
func (kern *Kernel) Equals(in Kernel) bool
```
Equals compares to kernels

#### type Layer

```go
type Layer interface {
}
```

Layer represents the general type of all layer types

#### type Net

```go
type Net struct {
}
```

Net is the basic type for Conv nets

#### func (*Net) AddLayer

```go
func (net *Net) AddLayer(lay Layer)
```
AddLayer adds another layer to the net

#### func (Net) Calculate

```go
func (net Net) Calculate() Volume
```
Calculate calcuates te output

#### type NormLayer

```go
type NormLayer struct {
	Layer
}
```

NormLayer is a normalisation layer

#### type PoolLayer

```go
type PoolLayer struct {
	Layer
}
```

PoolLayer will perform a downsampling operation along the spatial dimensions
(width, height), resulting in volume such as [16x16x12].

#### type ReluLayer

```go
type ReluLayer struct {
	Layer
}
```

ReluLayer will apply an elementwise activation function, such as the
max(0,x)max(0,x) thresholding at zero. This leaves the size of the volume
unchanged ([32x32x12]).

#### type Volume

```go
type Volume struct {
	Fields []mat64.Dense
}
```

Volume is a basic type to hold the layer's information

#### func  NewRNVolume

```go
func NewRNVolume(r, c, d int) *Volume
```
NewRNVolume generates a Volume of fixed size filled with zeros

#### func  NewRNVolumeRandom

```go
func NewRNVolumeRandom(r, c, d int) *Volume
```
NewRNVolumeRandom generates a Volume of fixed size filled with values between 0
and 1

#### func (*Volume) Apply

```go
func (vol *Volume) Apply(kern Kernel, strideR, strideC int)
```
Apply applys the given kernel to the whole volume, returnung a Volume with 1
depth

#### func (*Volume) Collumns

```go
func (vol *Volume) Collumns() int
```
Collumns of the Volume

#### func (*Volume) Depth

```go
func (vol *Volume) Depth() int
```
Depth of the Volume

#### func (*Volume) Dims

```go
func (vol *Volume) Dims() (int, int, int)
```
Dims returns the Dimensions of a Volume

#### func (*Volume) EqualSize

```go
func (vol *Volume) EqualSize(a Volume) bool
```
EqualSize checks if the size of two volumes are the same

#### func (*Volume) Equals

```go
func (vol *Volume) Equals(in Volume) bool
```
Equals compares the volume to another volume

#### func (*Volume) GetAt

```go
func (vol *Volume) GetAt(r, c, d int) float64
```
GetAt returns the element of the volume at a given position

#### func (Volume) Max

```go
func (vol Volume) Max() float64
```
Max returns the hightest number in a volume

#### func (*Volume) MulElem2

```go
func (vol *Volume) MulElem2(v1 Volume)
```
MulElem2 multiplies the volume with another volume element-wise

#### func (*Volume) PointReflect

```go
func (vol *Volume) PointReflect()
```
PointReflect calculates the pointreflection of a volume

#### func (*Volume) Print

```go
func (vol *Volume) Print()
```
Print prints the Volume to the console in a pretty format

#### func (*Volume) Reflect

```go
func (vol *Volume) Reflect()
```
Reflect calculates the reflectio of a volume (left-right)

#### func (*Volume) Rows

```go
func (vol *Volume) Rows() int
```
Rows of the Volume

#### func (*Volume) SetAll

```go
func (vol *Volume) SetAll(v Volume)
```
SetAll sets all values of the volume from another equal-sized volume

#### func (*Volume) SetAt

```go
func (vol *Volume) SetAt(r, c, d int, val float64)
```
SetAt sets the element of a volume at a given position

#### func (*Volume) SubVolumePadded

```go
func (vol *Volume) SubVolumePadded(cR, cC, r, c int) Volume
```
SubVolumePadded returns a part of the original Volume. cR and cC determine the
center of copying, r and c the size of the subvolume. If the size exceeds the
underlying volume the submodule is filled(padded with Zeros.
