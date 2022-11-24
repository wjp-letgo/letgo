package auth

import(
	"github.com/wjpxxx/letgo/lib"
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/auth/entity"
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"fmt"
)

//Auth
type Auth struct{
	Config *shopeeConfig.Config
}

//getBaseString
func (a *Auth)getBaseString(apiName string,ti int)string{
	return fmt.Sprintf("%d%s%d",a.Config.PartnerID,a.Config.GetApiPath(apiName),ti)
}
//AuthorizationURL
func (a *Auth)AuthorizationURL()string{
	ti:=lib.Time()
	baseString:=a.getBaseString("shop/auth_partner",ti)
	sign:=shopeeConfig.Sign(a.Config.PartnerKey,baseString)
	return fmt.Sprintf(
		"%s?partner_id=%d&timestamp=%d&redirect=%s&sign=%s",
		a.Config.GetApiURL("shop/auth_partner"),
		a.Config.PartnerID,
		ti,
		a.Config.RedirectURL,
		sign,
	)
}

//GetAccesstoken
func (a *Auth)GetAccesstoken(code string,shopID int64) entity.GetAccessTokenResult {
	method:="auth/token/get"
	params:=lib.InRow{
		"code":code,
		"partner_id":a.Config.PartnerID,
		"shop_id":shopID,
	}
	result:=entity.GetAccessTokenResult{}
	err:=a.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	result.ShopID=shopID
	return result
}

//RefreshAccessToken
func (a *Auth)RefreshAccessToken(shop commonentity.ShopInfo)entity.RefreshAccessTokenResult{
	method:="auth/access_token/get"
	params:=lib.InRow{
		"refresh_token":shop.RefreshToken,
		"partner_id":a.Config.PartnerID,
		"shop_id":shop.ShopID,
	}
	result:=entity.RefreshAccessTokenResult{}
	err:=a.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}