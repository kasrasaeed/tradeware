package rules

type And struct {
	rules []Rules
}

func NewAnd() *And {
	return &And{
		rules: make([]Rules, 0, 0),
	}
}

func (a *And) AddRule(rule ...Rules) {
	a.rules = append(a.rules, rule...)
}
func (a *And) Check() bool {
	for _, v := range a.rules {
		if !v.Check() {
			return false
		}
	}
	return true
}
