package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type KeltnerWidth struct {
	Indicator
}

func NewKeltnerWidth() IIndicator {
	return &KeltnerWidth{
		Indicator{
			Name: indicator_name.KELTNER_WIDTH,
			Settings: &KeltnerWidthSettings{
				EmaLength:  14,
				Multiplier: 2,
				AtrLength:  10,
				Source:     source_name.Close,
			},
			TimeFrame: timeframe.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type KeltnerWidthSettings struct {
	EmaLength  int
	Multiplier int
	AtrLength  int
	Source     source_name.SourceName
}
