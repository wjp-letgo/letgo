package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//GetTrackingNumberResult
type GetTrackingNumberResult struct{
	commonentity.Result
	Response TrackingNumberEntity `json:"response"`
}

//String
func(g GetTrackingNumberResult)String()string{
	return lib.ObjectToString(g)
}