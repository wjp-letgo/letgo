package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetChannelListResult
type GetChannelListResult struct {
	commonentity.Result
	Response GetChannelListResponse `json:"response"`
}

//String
func (g GetChannelListResult) String() string {
	return lib.ObjectToString(g)
}

//GetChannelListResponse
type GetChannelListResponse struct {
	LogisticsChannelList []LogisticsChannelEntity `json:"logistics_channel_list"`
	LogisticsDescription string                   `json:"logistics_description"`
	ForceEnabled         bool                     `json:"force_enabled"`
	MaskChannelID        int64                    `json:"mask_channel_id"`
}
