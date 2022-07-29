package position

import (
	"github.com/kasrasaeed/tradeware/v2/pkg/qty"
	"github.com/kasrasaeed/tradeware/v2/pkg/types/price"
	"github.com/kasrasaeed/tradeware/v2/pkg/types/side"
)

type Position struct {
	enteredPrice           price.Price
	stopLossPrice          price.Price
	takeProfitTriggerPrice price.Price
	side                   side.Side
	leverage               float64
	qty                    qty.Qty
	profitFactor           float64
	openTime               int64
	closeTime              int64
}

func New() *Position {
	return &Position{
		enteredPrice:           0,
		stopLossPrice:          0,
		takeProfitTriggerPrice: 0,
		leverage:               0,
		side:                   side.NotDefined,
		qty:                    0,
		openTime:               0,
		closeTime:              0,
	}
}

func (p *Position) SetLeverage(leverage float64) {
	p.leverage = leverage
}

func (p *Position) GetLeverage() float64 {
	return p.leverage
}

func (p *Position) SetEntryPrice(entryPrice price.Price) {
	p.enteredPrice = entryPrice
}

func (p *Position) GetEntryPrice() price.Price {
	return p.enteredPrice
}

func (p *Position) SetStopLossPrice(stopLossPrice price.Price) {
	p.stopLossPrice = stopLossPrice
}

func (p *Position) GetStopLossPrice() price.Price {
	return p.stopLossPrice
}

func (p *Position) SetProfitTriggerPrice(profitTriggerPrice price.Price) {
	p.takeProfitTriggerPrice = profitTriggerPrice
}

func (p *Position) GetProfitTriggerPrice() price.Price {
	return p.takeProfitTriggerPrice
}

func (p *Position) SetSide(side side.Side) {
	p.side = side
}

func (p *Position) GetSide() side.Side {
	return p.side
}

func (p *Position) SetQty(qty qty.Qty) {
	p.qty = qty
}

func (p *Position) GetQty() qty.Qty {
	return p.qty
}

func (p *Position) SetProfitFactor(profitFactor float64) {
	p.profitFactor = profitFactor
}

func (p *Position) GetProfitFactor() float64 {
	return p.profitFactor
}

func (p *Position) GetOpenTime() int64 {
	return p.openTime
}

func (p *Position) SetOpenTime(openTime int64) {
	p.openTime = openTime
}

func (p *Position) GetCloseTime() int64 {
	return p.closeTime
}

func (p *Position) SetCloseTime(closeTime int64) {
	p.closeTime = closeTime
}
