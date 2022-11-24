package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//DeleteGlobalModelResult
type DeleteGlobalModelResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response DeleteGlobalModelResultResponse `json:"response"`
}

//String
func(g DeleteGlobalModelResult)String()string{
	return lib.ObjectToString(g)
}
//DeleteGlobalModelResultResponse
type DeleteGlobalModelResultResponse struct{
	GlobalModelID int64 `json:"global_model_id"`
	Failures []FailureEntity `json:"failures"`
}

//String
func(g DeleteGlobalModelResultResponse)String()string{
	return lib.ObjectToString(g)
}