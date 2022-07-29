package indicator

import (
	"github.com/kasrasaeed/tradeware/v2/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v2/pkg/types/timeframe"
)

type Atr struct {
	Indicator
}

func NewAtr() IIndicator {
	return &Atr{
		Indicator{
			Name: indicator_name.ATR,
			Settings: &AtrSettings{
				AtrLength: 14,
			},
			TimeFrame: timeframe.ThreeMin,
			Value:     make([]float64, 0, 0),
		},
	}
}

type AtrSettings struct {
	AtrLength int
}
