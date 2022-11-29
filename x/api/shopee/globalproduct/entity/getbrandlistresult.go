package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetBrandListResult
type GetBrandListResult struct {
	commonentity.Result
	Response GetBrandListResultResponse `json:"response"`
	Warning  string                     `json:"warning"`
}

//String
func (g GetBrandListResult) String() string {
	return lib.ObjectToString(g)
}

//GetBrandListResultResponse
type GetBrandListResultResponse struct {
	BrandList   []BrandEntity `json:"brand_list"`
	HasNextPage bool          `json:"has_next_page"`
	NextOffset  int           `json:"next_offset"`
	IsMandatory bool          `json:"is_mandatory"`
	InputType   string        `json:"input_type"`
}
