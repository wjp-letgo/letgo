package config

import (
	"errors"
	"fmt"
	"strings"

	"github.com/wjp-letgo/letgo/encry"
	"github.com/wjp-letgo/letgo/httpclient"
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
	logisticsEntity "github.com/wjp-letgo/letgo/x/api/shopee/logistics/entity"
)

//Config
type Config struct{
	BaseURL string `json:"baseURL"`
	Version string `json:"version"`
	PartnerID int64 `json:"partner_id"`
	PartnerKey string `json:"partner_key"`
	RedirectURL string `json:"redirect_url"`
	shopInfo *commonentity.ShopInfo
}

//String
func (c *Config)String()string{
	return lib.ObjectToString(c)
}
//GetApiURL
func (c *Config)GetApiURL(apiPath string)string{
	return fmt.Sprintf("%s%s%s",c.BaseURL,c.Version,apiPath)
}

//GetApiPath
func (c *Config)GetApiPath(apiPath string)string{
	return fmt.Sprintf("%s%s",c.Version,apiPath)
}
//GetCommonParam
func (c *Config)GetCommonParam(method string) string{
	ti:=lib.Time()
	param:=lib.InRow{
		"partner_id":c.PartnerID,
		"timestamp":ti,
	}
	if c.shopInfo!=nil{
		param["access_token"]=c.shopInfo.AccessToken
		if strings.Index(method,"merchant")!=-1{
			param["merchant_id"]=c.shopInfo.ShopID
		}else{
			param["shop_id"]=c.shopInfo.ShopID
		}
		
		baseString:=fmt.Sprintf("%d%s%d%s%d",c.PartnerID,c.GetApiPath(method),ti,c.shopInfo.AccessToken,c.shopInfo.ShopID)
		param["sign"]=Sign(c.PartnerKey,baseString)
	}else{
		baseString:=fmt.Sprintf("%d%s%d",c.PartnerID,c.GetApiPath(method),ti)
		param["sign"]=Sign(c.PartnerKey,baseString)
	}
	return httpclient.HttpBuildQuery(param)
}

//HttpGet
func (c *Config)HttpGet(method string,data interface{},out interface{})error{
	return c.Http("GET", method,data,out)
}
//HttpPost
func (c *Config)HttpPost(method string,data interface{},out interface{})error{
	return c.Http("POST", method,data,out)
}

//HttpPostFile
func (c *Config)HttpPostFile(method string,data interface{},out interface{})error{
	return c.Http("POSTFILE", method,data,out)
}
//Http 请求
func (c *Config)Http(requestMethod,method string,data interface{},out interface{})error{
	query:=c.GetCommonParam(method)
	fullURL:=fmt.Sprintf("%s?%s",c.GetApiURL(method),query)
	ihttp:=httpclient.New().WithTimeOut(120)
	var result *httpclient.HttpResponse
	if requestMethod=="GET"{
		result=ihttp.Get(fullURL,data.(lib.InRow))
	}else if requestMethod=="POSTFILE"{
		result=ihttp.Post(fullURL,data.(lib.InRow))
	} else{
		result=ihttp.PostJson(fullURL,data)
	}
	if result.Err!=""{
		return errors.New(result.Err)
	}
	if result.Code!=200{
		return errors.New("请求失败")
	}
	if method=="logistics/download_shipping_document"{
		//下载快递单
		s:=lib.StringToObject(result.Body(),out)
		if !s{
			rs:=out.(*logisticsEntity.DownloadShippingDocumentResult)
			rs.File=result.BodyByte
		}
	}else{
		lib.StringToObject(result.Body(),out)
	}
	return nil
}

//SetShopInfo
func (c *Config)SetShopInfo(shopInfo *commonentity.ShopInfo)*Config{
	c.shopInfo=shopInfo
	return c
}

//New
func New(apiURL,version string,partnerID int64,partnerKey string,redirectURL string)*Config{
	return &Config{
		apiURL,
		version,
		partnerID,
		partnerKey,
		redirectURL,
		nil,
	}
}

//Sign
func Sign(partnerKey,baseString string)string{
	return encry.Hmac(baseString,partnerKey)
}