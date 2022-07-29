package rules

type FirstAboveSecond struct {
	firstValues    []float64
	secondValues   []float64
	firstIndexBar  int
	secondIndexBar int
}

func NewFirstAboveSecond(firstValues, SecondValues []float64, firstIndexBar, secondIndexBar int) *FirstAboveSecond {
	return &FirstAboveSecond{
		firstValues:    firstValues,
		secondValues:   SecondValues,
		firstIndexBar:  firstIndexBar,
		secondIndexBar: secondIndexBar,
	}
}

func (c *FirstAboveSecond) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.firstIndexBar] > secondValues[len(secondValues)-c.secondIndexBar] {
		return true
	}
	return false
}
