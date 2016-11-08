package main

type rNConvLayer struct {
	Filters []filter
}

func (l rNConvLayer) AddFilter(fil *filter) {
	l.Filters = append(l.Filters, *fil)
}

func (l rNConvLayer) calculate(vol rNVolume) *rNVolume {
	//TODO
	result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	//for k := range l.filters {

	//}
	return result
}

// Layer represents the general type of all layer types
type Layer interface {
}
