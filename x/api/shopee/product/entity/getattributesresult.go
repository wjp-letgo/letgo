package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetAttributesResult
type GetAttributesResult struct {
	commonentity.Result
	Response GetAttributesResultResponse `json:"response"`
	Warning  string                      `json:"warning"`
}

//String
func (g GetAttributesResult) String() string {
	return lib.ObjectToString(g)
}

//GetAttributesResultResponse
type GetAttributesResultResponse struct {
	AttributeList []AttributeEntity `json:"attribute_list"`
}

//String
func (g GetAttributesResultResponse) String() string {
	return lib.ObjectToString(g)
}
