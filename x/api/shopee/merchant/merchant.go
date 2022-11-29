package merchant

import (
	"github.com/wjp-letgo/letgo/lib"
	shopeeConfig "github.com/wjp-letgo/letgo/x/api/shopee/config"
	merchantEntity "github.com/wjp-letgo/letgo/x/api/shopee/merchant/entity"
)

//Merchant
type Merchant struct{
	Config *shopeeConfig.Config
}

//GetMerchantInfo
//@TitleUse this call to get information of merchant
//@Description https://open.shopee.com/documents?module=93&type=1&id=537&version=2
func (m *Merchant)GetMerchantInfo()merchantEntity.GetMerchantInfoResult{
	method:="merchant/get_merchant_info"
	result:=merchantEntity.GetMerchantInfoResult{}
	params:=lib.InRow{
	}
	err:=m.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetShopListByMerchant
//@Title Use this call to get shop_list bound to merchant_id.
//@Description https://open.shopee.com/documents?module=93&type=1&id=700&version=2
func (m *Merchant)GetShopListByMerchant(pageNo,pageSize int)merchantEntity.GetShopListByMerchantResult{
	method:="merchant/get_shop_list_by_merchant"
	result:=merchantEntity.GetShopListByMerchantResult{}
	params:=lib.InRow{
		"page_no":pageNo,
		"page_size":pageSize,
	}
	err:=m.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}