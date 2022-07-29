package rules

type FirstUnderSecond struct {
	firstValues    []float64
	secondValues   []float64
	firstIndexBar  int
	secondIndexBar int
}

func NewFirstUnderSecond(firstValues, SecondValues []float64, firstIndexBar, secondIndexBar int) *FirstUnderSecond {
	return &FirstUnderSecond{
		firstValues:    firstValues,
		secondValues:   SecondValues,
		firstIndexBar:  firstIndexBar,
		secondIndexBar: secondIndexBar,
	}
}

func (c *FirstUnderSecond) Check() bool {
	firstValues := c.firstValues
	secondValues := c.secondValues
	if firstValues[len(firstValues)-c.firstIndexBar] < secondValues[len(secondValues)-c.secondIndexBar] {
		return true
	}
	return false
}
