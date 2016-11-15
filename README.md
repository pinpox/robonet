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
func EqualVolDim(v1, v2 rNVolume) bool
```

#### func  NewRNVolume

```go
func NewRNVolume(r, c, d int) *rNVolume
```
NewRNVolume generates a rNVolume of fixed size filled with zeros

#### func  NewRNVolumeRandom

```go
func NewRNVolumeRandom(r, c, d int) *rNVolume
```
NewRNVolumeRandom generates a rNVolume of fixed size filled with values between
0 and 1

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
func (kern Kernel) Apply(in rNVolume) float64
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
func (kern *Kernel) SetAll(v rNVolume)
```

#### func (*Kernel) SetAt

```go
func (kern *Kernel) SetAt(r, c, d int, val float64)
```

#### func (*Kernel) Vol

```go
func (kern *Kernel) Vol() rNVolume
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
func (l *RNConvLayer) Calculate(vol rNVolume) rNVolume
```
Calculate applys all Kernels to a given Volume
