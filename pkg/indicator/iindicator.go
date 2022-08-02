package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type IIndicator interface {
	GetName() indicator_name.IndicatorName
	GetSettings() types.IndicatorSettings
	GetTimeFrame() timeframe.TimeFrame
	GetValue() []float64
	SetName(name indicator_name.IndicatorName)
	SetValue(value []float64)
	SetSettings(source *source_name.SourceName, values ...IndicatorSettingsAttr)
	SetTimeFrame(timeFrame timeframe.TimeFrame)
}
