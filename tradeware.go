package tradeware

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/indicator"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type TradeWare struct {
	timeFrames []timeframe.TimeFrame
	indicators []indicator.IIndicator
	Sources    []source_name.SourceName
	Symbol     string
}

func NewTradeWare(timeFrames []timeframe.TimeFrame, indicators []indicator.IIndicator, sources []source_name.SourceName, symbol string) *TradeWare {
	return &TradeWare{
		timeFrames: timeFrames,
		indicators: indicators,
		Sources:    sources,
		Symbol:     symbol,
	}
}
