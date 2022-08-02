package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type Adx struct {
	Indicator
}

func NewAdx() IIndicator {
	return &Adx{
		Indicator: Indicator{
			Name: indicator_name.ADX,
			Settings: &AdxSettings{
				AdxLength: 14,
			},
			TimeFrame: timeframe.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AdxSettings struct {
	AdxLength int
}
