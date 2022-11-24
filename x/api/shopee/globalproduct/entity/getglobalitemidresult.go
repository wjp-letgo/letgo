package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
)

//GetGlobalItemIDResult
type GetGlobalItemIDResult struct{
	commonentity.Result
	Warning string `json:"warning"`
	Response GetGlobalItemIDResultResponse `json:"response"`
}

//String
func(r GetGlobalItemIDResult)String()string{
	return lib.ObjectToString(r)
}
//GetGlobalItemIDResultResponse
type GetGlobalItemIDResultResponse struct{
	ItemIdMap []ItemIdMapEntity `json:"item_id_map"`
}

//String
func(r GetGlobalItemIDResultResponse)String()string{
	return lib.ObjectToString(r)
}