package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetShippingParameterResult
type GetShippingParameterResult struct{
	commonentity.Result
	Response GetShippingParameterResponse `json:"response"`
}

//String
func(g GetShippingParameterResult)String()string{
	return lib.ObjectToString(g)
}
//GetShippingParameterResponse
type GetShippingParameterResponse struct{
	InfoNeeded InfoNeededEntity `json:"info_needed"`
	Dropoff DropoffEntity `json:"dropoff"`
	Pickup PickupEntity `json:"pickup"`
}

//String
func(g GetShippingParameterResponse)String()string{
	return lib.ObjectToString(g)
}