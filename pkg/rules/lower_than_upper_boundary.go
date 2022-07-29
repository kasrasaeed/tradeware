package rules

type LowerThanUpperBoundary struct {
	values        []float64
	upperBoundary float64
	indexBar      int
}

func NewLowerThanUpperBoundary(values []float64, upperBoundary float64, indexBar int) *LowerThanUpperBoundary {
	return &LowerThanUpperBoundary{
		values:        values,
		upperBoundary: upperBoundary,
		indexBar:      indexBar,
	}
}

func (l *LowerThanUpperBoundary) Check() bool {
	latestValue := l.values[len(l.values)-l.indexBar]
	if latestValue < l.upperBoundary {
		return true
	}
	return false
}
