package globalproduct

import (
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	globalentity "github.com/wjpxxx/letgo/x/api/shopee/globalproduct/entity"
	"github.com/wjpxxx/letgo/lib"
)
const (
	NORMAL_BRAND BrandStatus =1
	PENDING_BRAND BrandStatus =2
)
//BrandStatus
type BrandStatus int

//GlobalProduct
type GlobalProduct struct{
	Config *shopeeConfig.Config
}

//GetCategory
//@Title Get global category. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=654&version=2
func (g *GlobalProduct)GetGlobalCategory(language string)globalentity.GetCategoryResult{
	method:="global_product/get_category"
	result:=globalentity.GetCategoryResult{}
	params:=lib.InRow{
		"language":language,
	}
	err:=g.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetAttributes
//@Title Get attributes. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=704&version=2
func (g *GlobalProduct)GetGlobalAttributes(language string,categoryID int64)globalentity.GetAttributesResult{
	method:="global_product/get_attributes"
	result:=globalentity.GetAttributesResult{}
	params:=lib.InRow{
		"language":language,
		"category_id":categoryID,
	}
	err:=g.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetBrandList
//@Title Use this call to get a list of brand. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=703&version=2
func (g *GlobalProduct)GetGlobalBrandList(offset,pageSize int,categoryID int64,status BrandStatus)globalentity.GetBrandListResult{
	method:="global_product/get_brand_list"
	result:=globalentity.GetBrandListResult{}
	params:=lib.InRow{
		"offset":offset,
		"page_size":pageSize,
		"category_id":categoryID,
		"status":status,
	}
	err:=g.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetItemLimit
//@Title Get global item upload control.Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=637&version=2
func (p *GlobalProduct)GetGlobalItemLimit(categoryID int64)globalentity.GetGlobalItemLimitResult{
	method:="global_product/get_global_item_limit"
	result:=globalentity.GetGlobalItemLimitResult{}
	err:=p.Config.HttpGet(method,nil,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetDtsLimit
//@Title Get category DTS limit. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=634&version=2
func (p *GlobalProduct)GetGlobalDtsLimit(categoryID int64)globalentity.GetDtsLimitResult{
	method:="global_product/get_dts_limit"
	result:=globalentity.GetDtsLimitResult{}
	params:=lib.InRow{
		"category_id":categoryID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetGlobalItemList
//@Title Use this call to get a list of items.
//@Description https://open.shopee.com/documents?module=89&type=1&id=614&version=2
func (p *GlobalProduct)GetGlobalItemList(offset string,pageSize int)globalentity.GetGlobalItemListResult{
	method:="global_product/get_global_item_list"
	params:=lib.InRow{
		"page_size":pageSize,
	}
	if offset!=""{
		params["offset"]=offset
	}
	result:=globalentity.GetGlobalItemListResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetGlobalItemInfo
//@Title Get global item info.Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926.
//@Description https://open.shopee.com/documents?module=90&type=1&id=644&version=2
func (p *GlobalProduct)GetGlobalItemInfo(globalItemIdList []int64)globalentity.GetGlobalItemInfoResult{
	method:="global_product/get_global_item_info"
	params:=lib.InRow{
		"global_item_id_list":globalItemIdList,
	}
	result:=globalentity.GetGlobalItemInfoResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//AddGlobalItem
//@Title Add a new item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=616&version=2
func (p *GlobalProduct)AddGlobalItem(item globalentity.AddItemRequestItemEntity)globalentity.AddGlobalItemResult{
	method:="global_product/add_global_item"
	result:=globalentity.AddGlobalItemResult{}
	err:=p.Config.HttpPost(method,item,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateGlobalItem
//@Title Update global item. Only for China mainland sellers using China Seller Centre(CNSC). More details in https://shopee.cn/cooperate/46/53/926
//@Description https://open.shopee.com/documents?module=90&type=1&id=620&version=2
func (p *GlobalProduct)UpdateGlobalItem(item globalentity.UpdateItemRequestItemEntity)globalentity.AddGlobalItemResult{
	method:="global_product/update_global_item"
	result:=globalentity.AddGlobalItemResult{}
	err:=p.Config.HttpPost(method,item,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DeleteGlobalItem
//@Title Delete global item. Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=621&version=2
func (p *GlobalProduct)DeleteGlobalItem(globalItemID int64)globalentity.DeleteGlobalItemResult{
	method:="global_product/delete_global_item"
	params:=lib.InRow{
		"global_item_id":globalItemID,
	}
	result:=globalentity.DeleteGlobalItemResult{}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//InitGlobalTierVariation
//@Title Initialize global product tier variation. Only for China mainland sellers using China Seller Centre(CNSC)
//@Description https://open.shopee.com/documents?module=90&type=1&id=635&version=2
func (p *GlobalProduct)InitGlobalTierVariation(globalItemID int64,tierVariation globalentity.TierVariationEntity,model globalentity.InitTierVariationModelEntity)globalentity.InitTierVariationResult{
	method:="global_product/init_tier_variation"
	result:=globalentity.InitTierVariationResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"tier_variation":tierVariation,
		"global_model":model,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateGlobalTierVariation
//@Title Update global product tier variation. Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=636&version=2
func (p *GlobalProduct)UpdateGlobalTierVariation(globalItemID int64,tierVariation []globalentity.TierVariationEntity)globalentity.UpdateTierVariationResult{
	method:="global_product/update_tier_variation"
	result:=globalentity.UpdateTierVariationResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"tier_variation":tierVariation,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//AddGlobalModel
//@Title Add model.
//@Description https://open.shopee.com/documents?module=90&type=1&id=643&version=2
func (p *GlobalProduct)AddGlobalModel(globalItemID int64,modelList []globalentity.InitTierVariationModelEntity)globalentity.AddGlobalModelResult{
	method:="global_product/add_global_model"
	result:=globalentity.AddGlobalModelResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"model_list":modelList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//UpdateGlobalModel
//@Title Update seller sku for model.
//@Description https://open.shopee.com/documents?module=90&type=1&id=645&version=2
func (p *GlobalProduct)UpdateGlobalModel(globalItemID int64,model []globalentity.UpdateModelEntity)globalentity.UpdateGlobalModelResult{
	method:="global_product/update_global_model"
	result:=globalentity.UpdateGlobalModelResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"global_model":model,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//DeleteGlobalModel
//@Title Delete item model.
//@Description https://open.shopee.com/documents?module=90&type=1&id=638&version=2
func (p *GlobalProduct)DeleteGlobalModel(globalItemID,globalModelID int64)globalentity.DeleteGlobalModelResult{
	method:="global_product/delete_global_model"
	result:=globalentity.DeleteGlobalModelResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"global_model_id":globalModelID,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetGlobalModelList
//@Title Get global model list. Only for China mainland sellers using China Seller Centre(CNSC)
//@Description https://open.shopee.com/documents?module=90&type=1&id=623&version=2
func (p *GlobalProduct)GetGlobalModelList(globalItemID int64)globalentity.GetModelListResult{
	method:="global_product/get_global_model_list"
	result:=globalentity.GetModelListResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//SupportGlobalSizeChart
//@Title Get category support size chart.
//@Description https://open.shopee.com/documents?module=90&type=1&id=632&version=2
func (p *GlobalProduct)SupportGlobalSizeChart(categoryID int64)globalentity.SupportGlobalSizeChartResult{
	method:="global_product/support_size_chart"
	result:=globalentity.SupportGlobalSizeChartResult{}
	params:=lib.InRow{
		"category_id":categoryID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateGlobalSizeChart
//@Title Update size chart image of item.
//@Description https://open.shopee.com/documents?module=90&type=1&id=625&version=2
func (p *GlobalProduct)UpdateGlobalSizeChart(globalItemID int64,sizeChart string)globalentity.UpdateGlobalSizeChartResult{
	method:="global_product/update_size_chart"
	result:=globalentity.UpdateGlobalSizeChartResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"size_chart":sizeChart,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//CreatePublishTask
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=639&version=2
func (p *GlobalProduct)CreatePublishTask(globalItemID,shopID int64,shopRegion string,item globalentity.CreatePublishTaskItemEntity)globalentity.CreatePublishTaskResult{
	method:="global_product/create_publish_task"
	result:=globalentity.CreatePublishTaskResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"shop_id":shopID,
		"shop_region":shopRegion,
		"item":item,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetPublishableShop
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=630&version=2
func (p *GlobalProduct)GetPublishableShop(globalItemID int64)globalentity.GetPublishableShopResult{
	method:="global_product/get_publishable_shop"
	result:=globalentity.GetPublishableShopResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}



//GetPublishTaskResult
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=627&version=2
func (p *GlobalProduct)GetPublishTaskResult(publishTaskID int64)globalentity.GetPublishTaskResult{
	method:="global_product/get_publish_task_result"
	result:=globalentity.GetPublishTaskResult{}
	params:=lib.InRow{
		"publish_task_id":publishTaskID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetPublishedList
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=633&version=2
func (p *GlobalProduct)GetPublishedList(globalItemID int64)globalentity.GetPublishedListResult{
	method:="global_product/get_published_list"
	result:=globalentity.GetPublishedListResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateGlobalPrice
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=642&version=2
func (p *GlobalProduct)UpdateGlobalPrice(globalItemID int64,priceList []globalentity.PriceEntity)globalentity.UpdateGlobalPriceResult{
	method:="global_product/update_price"
	result:=globalentity.UpdateGlobalPriceResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"price_list":priceList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//UpdateGlobalStock
//@Title Update stock.
//@Description https://open.shopee.com/documents?module=90&type=1&id=641&version=2
func (p *GlobalProduct)UpdateGlobalStock(globalItemID int64,stockList []globalentity.UpdateStockStockInfoEntity)globalentity.UpdateGlobalStockResult{
	method:="global_product/update_stock"
	result:=globalentity.UpdateGlobalStockResult{}
	params:=lib.InRow{
		"global_item_id":globalItemID,
		"stock_list":stockList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//SetSyncField
//@Title Update stock.
//@Description https://open.shopee.com/documents?module=90&type=1&id=656&version=2
func (p *GlobalProduct)SetSyncField(shopSyncList []globalentity.ShopSyncEntity)globalentity.SetSyncFieldResult{
	method:="global_product/set_sync_field"
	result:=globalentity.SetSyncFieldResult{}
	params:=lib.InRow{
		"shop_sync_list":shopSyncList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetGlobalItemID
//@Title Create publish task for global item.Only for China mainland sellers using China Seller Centre(CNSC).
//@Description https://open.shopee.com/documents?module=90&type=1&id=657&version=2
func (p *GlobalProduct)GetGlobalItemID(shopID int64,itemIdList []int64)globalentity.GetGlobalItemIDResult{
	method:="global_product/get_global_item_id"
	result:=globalentity.GetGlobalItemIDResult{}
	params:=lib.InRow{
		"shop_id":shopID,
		"item_id_list":itemIdList,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}