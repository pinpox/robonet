package main

type rNConvLayer struct {
	Filters []Filter
}

func (l rNConvLayer) AddFilter(fil *Filter) {
	l.Filters = append(l.Filters, *fil)
}

//Calculate applys all Filters to a given Volume
func (l rNConvLayer) calculate(vol rNVolume) *rNVolume {
	result := newRNVolume(vol.Height(), vol.Width(), vol.Depth())
	for _, v := range l.Filters {
		//TODO apply filtes
		vol = vol.Apply(v)
	}
	return result
}

// Layer represents the general type of all layer types
type Layer interface {
}
