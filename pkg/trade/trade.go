package trade

import "github.com/kasrasaeed/tradeware/v2/pkg/strategy"

type Trade struct {
	longStrategies  []*strategy.Strategy
	shortStrategies []*strategy.Strategy
}

func New() *Trade {
	return &Trade{
		longStrategies:  make([]*strategy.Strategy, 0),
		shortStrategies: make([]*strategy.Strategy, 0),
	}
}

func (t *Trade) AppendLongStrategy(longStrategies ...*strategy.Strategy) {
	t.longStrategies = append(t.longStrategies, longStrategies...)
}

func (t *Trade) AppendShortStrategy(shortStrategies ...*strategy.Strategy) {
	t.shortStrategies = append(t.shortStrategies, shortStrategies...)
}

func (t *Trade) GetLongStrategies() []*strategy.Strategy {
	return t.longStrategies
}

func (t *Trade) GetShortStrategies() []*strategy.Strategy {
	return t.shortStrategies
}
