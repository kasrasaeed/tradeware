package rules

type HigherThanUpperBoundary struct {
	values        []float64
	upperBoundary float64
	indexBar      int
}

func NewHigherThanUpperBoundary(values []float64, upperBoundary float64, indexBar int) *HigherThanUpperBoundary {
	return &HigherThanUpperBoundary{
		values:        values,
		upperBoundary: upperBoundary,
		indexBar:      indexBar,
	}
}
func (h *HigherThanUpperBoundary) Check() bool {
	latestValue := h.values[len(h.values)-h.indexBar]
	if latestValue > h.upperBoundary {
		return true
	}
	return false
}
