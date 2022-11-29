package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetItemListResult
type GetItemListResult struct {
	commonentity.Result
	Response GetItemListResultResponse `json:"response"`
	Warning  string                    `json:"warning"`
}

//String
func (g GetItemListResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemListResultResponse
type GetItemListResultResponse struct {
	Item        []ItemListEntity `json:"item"`
	TotalCount  int              `json:"total_count"`
	HasNextPage bool             `json:"has_next_page"`
	NextOffset  int              `json:"next_offset"`
}

//String
func (g GetItemListResultResponse) String() string {
	return lib.ObjectToString(g)
}
