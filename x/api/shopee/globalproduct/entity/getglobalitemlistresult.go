package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetItemListResult
type GetGlobalItemListResult struct {
	commonentity.Result
	Response GetGlobalItemListResultResponse `json:"response"`
	Warning  string                          `json:"warning"`
}

//String
func (g GetGlobalItemListResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemListResultResponse
type GetGlobalItemListResultResponse struct {
	GlobalItemList []GlobalItemListEntity `json:"global_item_list"`
	TotalCount     int                    `json:"total_count"`
	HasNextPage    bool                   `json:"has_next_page"`
	Offset         string                 `json:"offset"`
}

//String
func (g GetGlobalItemListResultResponse) String() string {
	return lib.ObjectToString(g)
}
