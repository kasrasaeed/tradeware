package kline

import (
	"github.com/kasrasaeed/tradeware/v2/pkg/types/timeframe"
)

type Kline struct {
	Open     float64
	High     float64
	Low      float64
	Close    float64
	Volume   float64
	OpenTime int64
	Interval timeframe.TimeFrame
}

func New(open, high, low, close, volume float64, OpenTime int64, interval timeframe.TimeFrame) *Kline {
	return &Kline{
		Open:     open,
		High:     high,
		Low:      low,
		Close:    close,
		Volume:   volume,
		OpenTime: OpenTime,
		Interval: interval,
	}
}
