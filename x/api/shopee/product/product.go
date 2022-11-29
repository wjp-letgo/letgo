package product

import (
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	shopeeConfig "github.com/wjp-letgo/letgo/x/api/shopee/config"
	"github.com/wjp-letgo/letgo/x/api/shopee/product/entity"
)

const (
	INT_TYPE AttributeType="INT_TYPE"
	STRING_TYPE AttributeType="STRING_TYPE"
	ENUM_TYPE AttributeType="ENUM_TYPE"
	FLOAT_TYPE AttributeType="FLOAT_TYPE"
	DATE_TYPE AttributeType="DATE_TYPE"
	TIMESTAMP_TYPE AttributeType="TIMESTAMP_TYPE"
	NORMAL ItemStatus="NORMAL"
	BANNED ItemStatus="BANNED"
	DELETED ItemStatus="DELETED"
	UNLIST ItemStatus="UNLIST"
	NORMAL_BRAND BrandStatus =1
	PENDING_BRAND BrandStatus =2
	REQUIRES_ATTRIBUTE AttributeStatus=1
	OPTIONAL_ATTRIBUTE AttributeStatus=2
)
//AttributeType
type AttributeType string

//ItemStatus
type ItemStatus string

//BrandStatus
type BrandStatus int
//AttributeStatus
type AttributeStatus int

//Product
type Product struct{
	Config *shopeeConfig.Config
}

//GetComment
//@Title Use this api to get comment by shop_id, item_id, or comment_id.
//@Description https://open.shopee.com/documents?module=89&type=1&id=562&version=2
func (p *Product)GetComment(itemID,commentID int64,cursor string,pageSize int)entity.GetCommentResult{
	method:="product/get_comment"
	params:=lib.InRow{
		"item_id":itemID,
		"comment_id":commentID,
		"cursor":cursor,
		"page_size":pageSize,
	}
	result:=entity.GetCommentResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//ReplyComment
//@Title Use this api to reply comments from buyers in batch.
//@Description https://open.shopee.com/documents?module=89&type=1&id=563&version=2
func (p *Product)ReplyComment(commentList []entity.ReplyCommentRequestCommentEntity)entity.ReplyCommentResult{
	method:="product/reply_comment"
	result:=entity.ReplyCommentResult{}
	err:=p.Config.HttpPost(method,commentList,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetItemBaseInfo
//@Title Use this api to get basic info of item by item_id list.
//@Description https://open.shopee.com/documents?module=89&type=1&id=612&version=2
func (p *Product)GetItemBaseInfo(itemIdList []int64)entity.GetItemBaseInfoResult{
	method:="product/get_item_base_info"
	params:=lib.InRow{
		"item_id_list":strings.Join(lib.Int64ArrayToArrayString(itemIdList),","),
	}
	result:=entity.GetItemBaseInfoResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetItemExtraInfo
//@Title Use this api to get extra info of item by item_id list.
//@Description https://open.shopee.com/documents?module=89&type=1&id=613&version=2
func (p *Product)GetItemExtraInfo(itemIdList []int64)entity.GetItemExtraInfoResult{
	method:="product/get_item_extra_info"
	params:=lib.InRow{
		"item_id_list":strings.Join(lib.Int64ArrayToArrayString(itemIdList),","),
	}
	result:=entity.GetItemExtraInfoResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetItemList
//@Title Use this call to get a list of items.
//@Description https://open.shopee.com/documents?module=89&type=1&id=614&version=2
func (p *Product)GetItemList(offset,pageSize,updateTimeFrom,updateTimeTo int,itemStatus ItemStatus)entity.GetItemListResult{
	method:="product/get_item_list"
	params:=lib.InRow{
		"offset":offset,
		"page_size":pageSize,
		"update_time_from":updateTimeFrom,
		"update_time_to":updateTimeTo,
		"item_status":itemStatus,
	}
	result:=entity.GetItemListResult{}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DeleteItem
//@Title Use this call to delete a product item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=615&version=2
func (p *Product)DeleteItem(itemID int64)entity.DeleteItemResult{
	method:="product/delete_item"
	params:=lib.InRow{
		"item_id":itemID,
	}
	result:=entity.DeleteItemResult{}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//AddItem
//@Title Add a new item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=616&version=2
func (p *Product)AddItem(item entity.AddItemRequestItemEntity)entity.AddItemResult{
	method:="product/add_item"
	result:=entity.AddItemResult{}
	err:=p.Config.HttpPost(method,item,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateItem
//@Title Update item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=617&version=2
func (p *Product)UpdateItem(item entity.UpdateItemRequestItemEntity)entity.UpdateItemResult{
	method:="product/update_item"
	result:=entity.UpdateItemResult{}
	err:=p.Config.HttpPost(method,item,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetModelList
//@Title Get model list of an item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=618&version=2
func (p *Product)GetModelList(itemID int64)entity.GetModelListResult{
	method:="product/get_model_list"
	result:=entity.GetModelListResult{}
	params:=lib.InRow{
		"item_id":itemID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}



//UpdateSizeChart
//@Title Update size chart image of item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=619&version=2
func (p *Product)UpdateSizeChart(itemID int64,sizeChart string)entity.UpdateSizeChartResult{
	method:="product/update_size_chart"
	result:=entity.UpdateSizeChartResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"size_chart":sizeChart,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//UnlistItem
//@Title Unlist item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=622&version=2
func (p *Product)UnlistItem(itemList []entity.UnlistItemItemListEntity)entity.UnlistItemResult{
	method:="product/unlist_item"
	result:=entity.UnlistItemResult{}
	err:=p.Config.HttpPost(method,itemList,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}



//BoostItem
//@Title Boost item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=624&version=2
func (p *Product)BoostItem(itemIdList []int64)entity.BoostItemResult{
	method:="product/boost_item"
	result:=entity.BoostItemResult{}
	params:=lib.InRow{
		"item_id_list":itemIdList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetBoostedList
//@Title Get boosted item list.
//@Description https://open.shopee.com/documents?module=89&type=1&id=626&version=2
func (p *Product)GetBoostedList(itemIdList []int64)entity.GetBoostedListResult{
	method:="product/get_boosted_list"
	result:=entity.GetBoostedListResult{}
	err:=p.Config.HttpGet(method,nil,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetDtsLimit
//@Title Get day to shipping limit.
//@Description https://open.shopee.com/documents?module=89&type=1&id=628&version=2
func (p *Product)GetDtsLimit(categoryID int64)entity.GetDtsLimitResult{
	method:="product/get_dts_limit"
	result:=entity.GetDtsLimitResult{}
	params:=lib.InRow{
		"category_id":categoryID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetItemLimit
//@Title Get item upload control.
//@Description https://open.shopee.com/documents?module=89&type=1&id=629&version=2
func (p *Product)GetItemLimit(categoryID int64)entity.GetItemLimitResult{
	method:="product/get_item_limit"
	result:=entity.GetItemLimitResult{}
	err:=p.Config.HttpGet(method,nil,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//SupportSizeChart
//@Title Get category support size chart.
//@Description https://open.shopee.com/documents?module=89&type=1&id=631&version=2
func (p *Product)SupportSizeChart(category int64)entity.SupportSizeChartResult{
	method:="product/support_size_chart"
	result:=entity.SupportSizeChartResult{}
	params:=lib.InRow{
		"category":category,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//InitTierVariation
//@Title Init item tier-variation struct.
//@Description https://open.shopee.com/documents?module=89&type=1&id=646&version=2
func (p *Product)InitTierVariation(itemID int64,tierVariation entity.TierVariationEntity,model entity.InitTierVariationModelEntity)entity.InitTierVariationResult{
	method:="product/init_tier_variation"
	result:=entity.InitTierVariationResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"tier_variation":tierVariation,
		"model":model,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateTierVariation
//@Title Update item tier-variation struct.
//@Description https://open.shopee.com/documents?module=89&type=1&id=647&version=2
func (p *Product)UpdateTierVariation(itemID int64,tierVariation []entity.TierVariationEntity)entity.UpdateTierVariationResult{
	method:="product/update_tier_variation"
	result:=entity.UpdateTierVariationResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"tier_variation":tierVariation,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateModel
//@Title Update seller sku for model.
//@Description https://open.shopee.com/documents?module=89&type=1&id=648&version=2
func (p *Product)UpdateModel(itemID int64,model []entity.UpdateModelEntity)entity.UpdateModelResult{
	method:="product/update_model"
	result:=entity.UpdateModelResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"model":model,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//AddModel
//@Title Add model.
//@Description https://open.shopee.com/documents?module=89&type=1&id=649&version=2
func (p *Product)AddModel(itemID int64,modelList []entity.InitTierVariationModelEntity)entity.AddModelResult{
	method:="product/add_model"
	result:=entity.AddModelResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"model_list":modelList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DeleteModel
//@Title Delete item model.
//@Description https://open.shopee.com/documents?module=89&type=1&id=650&version=2
func (p *Product)DeleteModel(itemID,modelID int64)entity.DeleteModelResult{
	method:="product/delete_model"
	result:=entity.DeleteModelResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"model_id":modelID,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdatePrice
//@Title Update price.
//@Description https://open.shopee.com/documents?module=89&type=1&id=651&version=2
func (p *Product)UpdatePrice(itemID int64,priceList []entity.UpdatePricePriceInfoEntity)entity.UpdatePriceResult{
	method:="product/update_price"
	result:=entity.UpdatePriceResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"price_list":priceList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}
//UpdateStock
//@Title Update stock.
//@Description https://open.shopee.com/documents?module=89&type=1&id=652&version=2
func (p *Product)UpdateStock(itemID int64,stockList []entity.UpdateStockStockInfoEntity)entity.UpdateStockResult{
	method:="product/update_stock"
	result:=entity.UpdateStockResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"stock_list":stockList,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}
//GetCategory
//@Title Get category.
//@Description https://open.shopee.com/documents?module=89&type=1&id=653&version=2
func (p *Product)GetCategory(language string)entity.GetCategoryResult{
	method:="product/get_category"
	result:=entity.GetCategoryResult{}
	params:=lib.InRow{
		"language":language,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetAttributes
//@Title Get attributes.
//@Description https://open.shopee.com/documents?module=89&type=1&id=655&version=2
func (p *Product)GetAttributes(language string,categoryID int64)entity.GetAttributesResult{
	method:="product/get_attributes"
	result:=entity.GetAttributesResult{}
	params:=lib.InRow{
		"language":language,
		"category_id":categoryID,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetBrandList
//@Title Use this call to get a list of brand.
//@Description https://open.shopee.com/documents?module=89&type=1&id=684&version=2
func (p *Product)GetBrandList(offset,pageSize int,categoryID int64,status BrandStatus)entity.GetBrandListResult{
	method:="product/get_brand_list"
	result:=entity.GetBrandListResult{}
	params:=lib.InRow{
		"offset":offset,
		"page_size":pageSize,
		"category_id":categoryID,
		"status":status,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//CategoryRecommend
//@Title Recommend category by item name.
//@Description https://open.shopee.com/documents?module=89&type=1&id=702&version=2
func (p *Product)CategoryRecommend(itemName string)entity.CategoryRecommendResult{
	method:="product/category_recommend"
	result:=entity.CategoryRecommendResult{}
	params:=lib.InRow{
		"item_name":itemName,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}
//GetItemPromotion
//@Title Get item promotion info.
//@Description https://open.shopee.com/documents?module=89&type=1&id=661&version=2
func (p *Product)GetItemPromotion(itemIdList []int64)entity.GetItemPromotionResult{
	method:="product/get_item_promotion"
	result:=entity.GetItemPromotionResult{}
	params:=lib.InRow{
		"item_name":itemIdList,
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateSipItemPrice
//@Title Update sip item price.
//@Description https://open.shopee.com/documents?module=89&type=1&id=662&version=2
func (p *Product)UpdateSipItemPrice(itemID int64,sipItemPrice []entity.SipItemPriceEntity)entity.UpdateSipItemPriceResult{
	method:="product/update_sip_item_price"
	result:=entity.UpdateSipItemPriceResult{}
	params:=lib.InRow{
		"item_id":itemID,
		"sip_item_price":sipItemPrice,
	}
	err:=p.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//SearchItem
//@Title Use this call to search item.
//@Description https://open.shopee.com/documents?module=89&type=1&id=701&version=2
func (p *Product)SearchItem(offset string,pageSize int,itemName string,attributeStatus AttributeStatus)entity.SearchItemResult{
	method:="product/search_item"
	result:=entity.SearchItemResult{}
	params:=lib.InRow{
		"page_size":pageSize,
	}
	if offset!=""{
		params["offset"]=offset
	}
	if itemName!=""{
		params["item_name"]=itemName
	}
	if attributeStatus>0{
		params["attribute_status"]=attributeStatus
	}
	err:=p.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}