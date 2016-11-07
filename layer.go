package main

type rNConvLayer struct {
	Filters []filter
}

func (l rNConvLayer) AddFilter(fil *filter) {
	//TODO
}

func (l rNConvLayer) calculate(vol rNVolume) *rNVolume {
	result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	//for k := range l.filters {

	//}
	return result
}

// Layer represents the general type of all layer types
type Layer interface {
}
