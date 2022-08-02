package strategy

import "github.com/kasrasaeed/tradeware/v3/pkg/rules"

type IStrategy interface {
	AddEntryRule(er rules.Rules)

	AddTakeProfitRule(tpr rules.Rules)

	AddStopLossRule(slr rules.Rules)

	UpdateEntryRule(er rules.Rules)

	UpdateTakeProfitRule(tp rules.Rules)

	UpdateStopLossRule(sl rules.Rules)

	CheckEntry() bool

	CheckTakeProfit() bool

	CheckSlHit() bool
}
