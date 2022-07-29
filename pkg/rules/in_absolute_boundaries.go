package rules

type InAbsoluteBoundaries struct {
	upperBoundary float64
	lowerBoundary float64
	values        []float64
}

func NewInAbsoluteBoundaries(upperBoundary, lowerBoundary float64, values []float64) *InAbsoluteBoundaries {
	return &InAbsoluteBoundaries{
		upperBoundary: upperBoundary,
		lowerBoundary: lowerBoundary,
		values:        values,
	}
}

func (i *InAbsoluteBoundaries) Check() bool {
	var startIndex int

	for index := len(i.values) - 1; index >= 0; index-- {
		if i.values[index] < i.upperBoundary && i.values[index] >= i.lowerBoundary && i.values[index-1] < i.lowerBoundary {
			startIndex = index
			break
		}
	}
	if startIndex == 0 {
		return false
	}
	for index := len(i.values) - 1; index >= startIndex; index-- {
		if !(i.values[index] < i.upperBoundary && i.values[index] > i.lowerBoundary) {
			return false
		}
	}
	return true
}
