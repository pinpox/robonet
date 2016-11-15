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

#### func  Odd3Dim

```go
func Odd3Dim(i1, i2, i3 int) bool
```
Odd3Dim checks if the height and width are odd

#### func  SigmoidFast

```go
func SigmoidFast(x float64) float64
```
SigmoidFast calcultes the value for activation using a fast sigmoid
approximation

#### type Kernel

```go
type Kernel struct {
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

#### func (Kernel) Dims

```go
func (kern Kernel) Dims() (int, int, int)
```
Dims returns the size of the kernel

#### func (*Kernel) Equals

```go
func (kern *Kernel) Equals(in Kernel) bool
```

#### func (*Kernel) GetAt

```go
func (kern *Kernel) GetAt(r, c, d int) float64
```

#### func (*Kernel) PointReflect

```go
func (kern *Kernel) PointReflect()
```

#### func (Kernel) Print

```go
func (kern Kernel) Print()
```
Print shows show the kernel's matrix string representation

#### func (*Kernel) Reflect

```go
func (kern *Kernel) Reflect()
```

#### func (*Kernel) SetAll

```go
func (kern *Kernel) SetAll(v Volume)
```

#### func (*Kernel) SetAt

```go
func (kern *Kernel) SetAt(r, c, d int, val float64)
```

#### func (*Kernel) Vol

```go
func (kern *Kernel) Vol() Volume
```

#### type Layer

```go
type Layer interface {
}
```

Layer represents the general type of all layer types

#### type RNConvLayer

```go
type RNConvLayer struct {
	Kernels []Kernel
}
```

RNConvLayer basic type for a convolutional layer

#### func (*RNConvLayer) AddKernel

```go
func (l *RNConvLayer) AddKernel(fil Kernel)
```
AddKernel adds a kernel to a layer

#### func (*RNConvLayer) Calculate

```go
func (l *RNConvLayer) Calculate(vol Volume) Volume
```
Calculate applys all Kernels to a given Volume

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
func (vol *Volume) Apply(f Kernel)
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

#### func (*Volume) EqualSize

```go
func (vol *Volume) EqualSize(a Volume) bool
```
EqualSize checks if the size of two volumes are the same

#### func (*Volume) Equals

```go
func (vol *Volume) Equals(in Volume) bool
```

#### func (*Volume) GetAt

```go
func (vol *Volume) GetAt(r, c, d int) float64
```

#### func (Volume) Max

```go
func (vol Volume) Max() float64
```

#### func (*Volume) MulElem2

```go
func (vol *Volume) MulElem2(v1 Volume)
```

#### func (*Volume) PointReflect

```go
func (vol *Volume) PointReflect()
```

#### func (*Volume) Print

```go
func (vol *Volume) Print()
```

#### func (*Volume) Reflect

```go
func (vol *Volume) Reflect()
```

#### func (*Volume) Rows

```go
func (vol *Volume) Rows() int
```
Rows of the Volume

#### func (*Volume) SetAt

```go
func (vol *Volume) SetAt(r, c, d int, val float64)
```

#### func (*Volume) SubVolumePadded

```go
func (vol *Volume) SubVolumePadded(cR, cC, r, c int) Volume
```
SubVolumePadded returns a part of the original Volume. cR and cC determine the
center of copying, r and c the size of the subvolume. If the size exceeds the
underlying volume the submodule is filled(padded with Zeros.
