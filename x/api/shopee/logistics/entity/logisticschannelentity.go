package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//LogisticsChannelEntity
type LogisticsChannelEntity struct{
	LogisticsChannelID int64 `json:"logistics_channel_id"`
	Preferred bool `json:"preferred"`
	LogisticsChannelName string `json:"logistics_channel_name"`
	CodEnabled bool `json:"cod_enabled"`
	Enabled bool `json:"enabled"`
	FeeType string `json:"fee_type"`
	SizeList []SizeEntity `json:"size_list"`
	WeightLimit WeightLimitEntity `json:"weight_limit"`
	ItemMaxDimension ItemMaxDimensionEntity `json:"item_max_dimension"`
	VolumeLimit VolumeLimitEntity `json:"volume_limit"`
}

//String
func(l LogisticsChannelEntity)String()string{
	return lib.ObjectToString(l)
}