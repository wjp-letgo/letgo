package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetModelListResult
type GetModelListResult struct{
	commonentity.Result
	Response GetModelListResultResponse `json:"response"`
	Warning string `json:"warning"`
}

//String
func(g GetModelListResult)String()string{
	return lib.ObjectToString(g)
}
//GetModelListResultResponse
type GetModelListResultResponse struct{
	TierVariation []TierVariationEntity `json:"tier_variation"`
	Model []ModelEntity `json:"model"`
}

//String
func(g GetModelListResultResponse)String()string{
	return lib.ObjectToString(g)
}