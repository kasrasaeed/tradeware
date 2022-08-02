package strategy

import (
	"github.com/kasrasaeed/tradeware/v2/pkg/rules"
)

type Strategy struct {
	name            string
	entryRules      []rules.Rules
	takeProfitRules []rules.Rules
	stopLossRules   []rules.Rules
}

func New(name string, entryRules, takeProfitRules, stopLossRules []rules.Rules) *Strategy {
	return &Strategy{
		name:            name,
		entryRules:      entryRules,
		takeProfitRules: takeProfitRules,
		stopLossRules:   stopLossRules,
	}
}

func (s *Strategy) AddEntryRule(er rules.Rules) {
	s.entryRules = append(s.entryRules, er)
}

func (s *Strategy) AddTakeProfitRule(tpr rules.Rules) {
	s.takeProfitRules = append(s.takeProfitRules, tpr)
}

func (s *Strategy) AddStopLossRule(slr rules.Rules) {
	s.stopLossRules = append(s.stopLossRules, slr)
}

func (s *Strategy) UpdateEntryRules(er ...rules.Rules) {
	s.entryRules = make([]rules.Rules, 0)
	s.entryRules = append(s.entryRules, er...)
}

func (s *Strategy) UpdateTakeProfitRules(tp ...rules.Rules) {
	s.takeProfitRules = make([]rules.Rules, 0)
	s.takeProfitRules = append(s.takeProfitRules, tp...)
}

func (s *Strategy) UpdateStopLossRules(sl ...rules.Rules) {
	s.stopLossRules = make([]rules.Rules, 0)
	s.stopLossRules = append(s.stopLossRules, sl...)
}

func (s *Strategy) CheckEntry() bool {
	for _, v := range s.entryRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckTakeProfit() bool {
	for _, v := range s.takeProfitRules {
		if !v.Check() {
			return false
		}
	}
	return true
}

func (s *Strategy) CheckSlHit() bool {
	for _, v := range s.stopLossRules {
		if !v.Check() {
			return false
		}
	}
	return true
}
