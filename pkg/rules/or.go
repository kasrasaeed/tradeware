package rules

type Or struct {
	rules []Rules
}

func NewOr() *Or {
	return &Or{
		rules: make([]Rules, 0, 0),
	}
}

func (o *Or) AddRule(rule ...Rules) {
	o.rules = append(o.rules, rule...)
}
func (o *Or) Check() bool {
	for _, v := range o.rules {
		if v.Check() {
			return true
		}
	}
	return false
}
