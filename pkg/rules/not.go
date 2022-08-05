package rules

type Not struct {
	rule Rules
}

func NewNot(rule Rules) *Not {
	return &Not{
		rule: rule,
	}
}

func (n *Not) Check() bool {
	return !n.rule.Check()
}
