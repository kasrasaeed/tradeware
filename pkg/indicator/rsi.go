package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type Rsi struct {
	Indicator
}

func NewRsi() IIndicator {
	return &Rsi{
		Indicator: Indicator{
			Name: indicator_name.RSI,
			Settings: &RsiSettings{
				RsiLength:  14,
				SourceName: source_name.Close,
			},
			TimeFrame: timeframe.FiveMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type RsiSettings struct {
	RsiLength  int
	SourceName source_name.SourceName
}
