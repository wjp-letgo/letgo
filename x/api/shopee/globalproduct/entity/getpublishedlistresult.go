package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetPublishedListResult
type GetPublishedListResult struct {
	commonentity.Result
	Response GetPublishedListResultResponse `json:"response"`
	Warning  string                         `json:"warning"`
}

//String
func (g GetPublishedListResult) String() string {
	return lib.ObjectToString(g)
}

//GetPublishedListResultResponse
type GetPublishedListResultResponse struct {
	PublishedItem []PublishedItemEntity `json:"published_item"`
}

//String
func (g GetPublishedListResultResponse) String() string {
	return lib.ObjectToString(g)
}
