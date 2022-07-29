package rules

type CrossUp struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCrossUp(firstValues, SecondValues []float64, indexBar int) *CrossUp {
	return &CrossUp{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (c *CrossUp) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.indexBar] > secondValues[len(secondValues)-c.indexBar] &&
		firstValues[len(firstValues)-c.indexBar-1] < secondValues[len(secondValues)-c.indexBar-1] {
		return true
	}
	return false
}
