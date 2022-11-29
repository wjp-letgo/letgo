package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//SearchItemResult
type SearchItemResult struct {
	commonentity.Result
	Response SearchItemResultResponse `json:"response"`
	Warning  []string                 `json:"warning"`
}

//String
func (r SearchItemResult) String() string {
	return lib.ObjectToString(r)
}

//SearchItemResultResponse
type SearchItemResultResponse struct {
	ItemIdList []int64 `json:"item_id_list"`
	TotalCount int     `json:"total_count"`
	NextOffset string  `json:"next_offset"`
}

//String
func (r SearchItemResultResponse) String() string {
	return lib.ObjectToString(r)
}
