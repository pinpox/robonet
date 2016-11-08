package robonet

//RNConvLayer basic type for a convolutional layer
type RNConvLayer struct {
	Filters []Filter
}

//AddFilter adds a filter to a layer
func (l RNConvLayer) AddFilter(fil *Filter) {
	l.Filters = append(l.Filters, *fil)
}

//Calculate applys all Filters to a given Volume
func (l RNConvLayer) Calculate(vol rNVolume) rNVolume {
	//result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	for _, v := range l.Filters {
		vol = vol.Apply(v)
	}
	return vol
}

// Layer represents the general type of all layer types
type Layer interface {
}
