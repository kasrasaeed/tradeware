package indicator

import (
	"github.com/kasrasaeed/tradeware/v2/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v2/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v2/pkg/types/timeframe"
)

type Ema struct {
	Indicator
}

func NewEma() IIndicator {
	return &Ema{
		Indicator: Indicator{
			Name: indicator_name.EMA,
			Settings: &EmaSettings{
				EmaLength: 14,
				Source:    source_name.Close,
			},
			TimeFrame: timeframe.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type EmaSettings struct {
	EmaLength int
	Source    source_name.SourceName
}
