# robonet
--
    import "gitlab.com/binaryplease/robonet"


## Usage

#### func  Equal3Dim

```go
func Equal3Dim(e1, e2, e3, i1, i2, i3 int) bool
```
Equal3Dim checks if the size of two volumes are the same

#### func  NewRNVolume

```go
func NewRNVolume(w, h, d int) *rNVolume
```
NewRNVolume generates a rNVolume of fixed size filled with zeros

#### func  NewRNVolumeRandom

```go
func NewRNVolumeRandom(w, h, d int) *rNVolume
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

#### type Filter

```go
type Filter struct {
}
```

Filter represets a basic conv filter

#### func  NewFilter

```go
func NewFilter(height int, width, depth int) *Filter
```
NewFilter creates a new filter initialized with zeros

#### func  NewFilterRandom

```go
func NewFilterRandom(height int, width, depth int) *Filter
```
NewFilterRandom creates a new filter initialized with random values

#### func (Filter) Apply

```go
func (f Filter) Apply(in rNVolume) float64
```
Apply applys the filter to a equally sized chunk of a volume Only filters of the
same size as the volume can be applied

#### func (Filter) Dims

```go
func (f Filter) Dims() (int, int, int)
```
Dims returns the size of the filter

#### func (Filter) Print

```go
func (f Filter) Print()
```
Print shows show the filter's matrix string representation

#### type Layer

```go
type Layer interface {
}
```

Layer represents the general type of all layer types

#### type RNConvLayer

```go
type RNConvLayer struct {
	Filters []Filter
}
```

RNConvLayer basic type for a convolutional layer

#### func (RNConvLayer) AddFilter

```go
func (l RNConvLayer) AddFilter(fil *Filter)
```
AddFilter adds a filter to a layer

#### func (RNConvLayer) Calculate

```go
func (l RNConvLayer) Calculate(vol rNVolume) rNVolume
```
Calculate applys all Filters to a given Volume
