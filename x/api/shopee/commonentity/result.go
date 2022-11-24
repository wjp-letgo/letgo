package commonentity

import(
	"github.com/wjpxxx/letgo/lib"
)

//Result
type Result struct{
	Error string `json:"error"`
	Message string `json:"message"`
	RequestID string `json:"request_id"`
}
//String
func(r Result)String()string{
	return lib.ObjectToString(r)
}

//ShopInfo
type ShopInfo struct{
	RefreshToken string `json:"refresh_token"`
	AccessToken string `json:"access_token"`
	ExpireIn int `json:"expire_in"`
	ShopID int64 `json:"shop_id"`
}

//String
func(s ShopInfo)String()string{
	return lib.ObjectToString(s)
}

//NewShop
func NewShop(shopID int64,expireIn int,accessToken,refreshToken string)*ShopInfo{
	return &ShopInfo{
		ShopID:shopID,
		ExpireIn:expireIn,
		AccessToken:accessToken,
		RefreshToken:refreshToken,
	}
}