package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type SuperTrend struct {
	Indicator
}

func NewSuperTrend() IIndicator {
	return &SuperTrend{
		Indicator{
			Name: indicator_name.SUPER_TREND,
			Settings: &SuperTrendSettings{
				AtrLength:  10,
				Multiplier: 4.4,
				Source:     source_name.Hlcc4,
			},
			TimeFrame: timeframe.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type SuperTrendSettings struct {
	AtrLength  int
	Multiplier float64
	Source     source_name.SourceName
}
