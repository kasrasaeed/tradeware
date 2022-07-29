package rules

type CrossDown struct {
	firstValues  []float64
	secondValues []float64
	indexBar     int
}

func NewCrossDown(firstValues, SecondValues []float64, indexBar int) *CrossDown {
	return &CrossDown{
		firstValues:  firstValues,
		secondValues: SecondValues,
		indexBar:     indexBar,
	}
}

func (c *CrossDown) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.indexBar] < secondValues[len(secondValues)-c.indexBar] &&
		firstValues[len(firstValues)-(c.indexBar-1)] > secondValues[len(secondValues)-(c.indexBar-1)] {
		return true
	}
	return false
}
