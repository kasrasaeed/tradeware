package indicator

import (
	"github.com/kasrasaeed/tradeware/v3/pkg/types"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/indicator_settings_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/source_name"
	"github.com/kasrasaeed/tradeware/v3/pkg/types/timeframe"
)

type Indicator struct {
	Name      indicator_name.IndicatorName
	Settings  types.IndicatorSettings
	TimeFrame timeframe.TimeFrame
	Value     []float64
}

func (i *Indicator) GetValue() []float64 {
	return i.Value
}

func (i *Indicator) GetName() indicator_name.IndicatorName {
	return i.Name
}

func (i *Indicator) GetSettings() types.IndicatorSettings {
	switch i.Settings.(type) {
	case *SmaSettings:
		return i.Settings.(*SmaSettings)
	case *DemaSettings:
		return i.Settings.(*DemaSettings)
	case *RsiSettings:
		return i.Settings.(*RsiSettings)
	case *AdxSettings:
		return i.Settings.(*AdxSettings)
	case *EmaSettings:
		return i.Settings.(*EmaSettings)
	case *AtrSettings:
		return i.Settings.(*AtrSettings)
	case *KeltnerWidthSettings:
		return i.Settings.(*KeltnerWidthSettings)
	case *SuperTrendSettings:
		return i.Settings.(*SuperTrendSettings)
	default:
		return i.Settings
	}
}

func (i *Indicator) GetTimeFrame() timeframe.TimeFrame {
	return i.TimeFrame
}

func (i *Indicator) SetValue(value []float64) {
	i.Value = value
}

func (i *Indicator) SetName(name indicator_name.IndicatorName) {
	i.Name = name
}

func (i *Indicator) SetSettings(source *source_name.SourceName, attrs ...IndicatorSettingsAttr) {
	switch i.Settings.(type) {
	case *SmaSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.SmaLength {
				i.Settings.(*SmaSettings).SmaLength = int(v.Value)
			}
		}
	case *DemaSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.DemaLength {
				i.Settings.(*DemaSettings).DemaLength = int(v.Value)
			}
		}
		i.Settings.(*DemaSettings).SourceName = *source
	case *RsiSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.RsiLength {
				i.Settings.(*RsiSettings).RsiLength = (int)(v.Value)
			}
		}
		i.Settings.(*RsiSettings).SourceName = *source
	case *AdxSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.AdxLength {
				i.Settings.(*AdxSettings).AdxLength = (int)(v.Value)
			}
		}
	case *EmaSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.EmaLength {
				i.Settings.(*EmaSettings).EmaLength = (int)(v.Value)
			}
		}
	case *AtrSettings:
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.AtrLength {
				i.Settings.(*AtrSettings).AtrLength = (int)(v.Value)
			}
		}
	case *KeltnerWidthSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.EmaLength {
				i.Settings.(*KeltnerWidthSettings).EmaLength = (int)(v.Value)
			} else if v.Attr == indicator_settings_name.AtrLength {
				i.Settings.(*KeltnerWidthSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == indicator_settings_name.Multiplier {
				i.Settings.(*KeltnerWidthSettings).Multiplier = (int)(v.Value)
			}
		}
		i.Settings.(*KeltnerWidthSettings).Source = *source

	case *SuperTrendSettings:
		if source == nil {
			panic((any)("source can not be nil"))
		}
		for _, v := range attrs {
			if v.Attr == indicator_settings_name.AtrLength {
				i.Settings.(*SuperTrendSettings).AtrLength = (int)(v.Value)
			} else if v.Attr == indicator_settings_name.Multiplier {
				i.Settings.(*SuperTrendSettings).Multiplier = v.Value
			}
		}
		i.Settings.(*SuperTrendSettings).Source = *source

	}
}

func (i *Indicator) SetTimeFrame(timeFrame timeframe.TimeFrame) {
	i.TimeFrame = timeFrame
}
