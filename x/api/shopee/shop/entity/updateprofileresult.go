package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//UpdateProfileResult
type UpdateProfileResult struct {
	commonentity.Result
	Response UpdateProfileResultResponse `json:"response"`
}

//String
func (g UpdateProfileResult) String() string {
	return lib.ObjectToString(g)
}

//UpdateProfileResultResponse
type UpdateProfileResultResponse struct {
	ShopLogo    string `json:"shop_logo"`
	Description string `json:"description"`
	ShopName    string `json:"shop_name"`
}

//String
func (g UpdateProfileResultResponse) String() string {
	return lib.ObjectToString(g)
}
