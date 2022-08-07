package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type Sma struct {
	Indicator
}

func NewSma() IIndicator {
	return &Sma{
		Indicator{
			Name: indicator_name.SMA,
			Settings: &SmaSettings{
				SmaLength: 20,
			},
			TimeFrame: timeframe.FifteenMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type SmaSettings struct {
	SmaLength int
}
