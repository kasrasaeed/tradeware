package tradeware

type TradeWare struct {
	timeFrames []TimeFrame
	indicators []IIndicator
	Sources    []Source
	Symbol     string
}

func NewTradeWare(timeFrames []TimeFrame, indicators []IIndicator, sources []Source, symbol string) *TradeWare {
	return &TradeWare{
		timeFrames: timeFrames,
		indicators: indicators,
		Sources:    sources,
		Symbol:     symbol,
	}
}
