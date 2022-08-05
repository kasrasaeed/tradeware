package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type Dema struct {
	Indicator
}

func NewDema() IIndicator {
	return &Dema{
		Indicator{
			Name: indicator_name.DEMA,
			Settings: &DemaSettings{
				DemaLength: 100,
				SourceName: source_name.Close,
			},
			TimeFrame: timeframe.TwoHour,
			Value:     make([]float64, 0, 0),
		},
	}
}

type DemaSettings struct {
	DemaLength int
	SourceName source_name.SourceName
}
