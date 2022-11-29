package shopee

import (
	"github.com/wjp-letgo/letgo/x/api/shopee/auth"
	authEntity "github.com/wjp-letgo/letgo/x/api/shopee/auth/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
	shopeeConfig "github.com/wjp-letgo/letgo/x/api/shopee/config"
	"github.com/wjp-letgo/letgo/x/api/shopee/globalproduct"
	globalProductEntity "github.com/wjp-letgo/letgo/x/api/shopee/globalproduct/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/logistics"
	logisticsEntity "github.com/wjp-letgo/letgo/x/api/shopee/logistics/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/mediaspace"
	mediaspaceEntity "github.com/wjp-letgo/letgo/x/api/shopee/mediaspace/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/merchant"
	merchantEntity "github.com/wjp-letgo/letgo/x/api/shopee/merchant/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/order"
	orderEntity "github.com/wjp-letgo/letgo/x/api/shopee/order/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/product"
	productEntity "github.com/wjp-letgo/letgo/x/api/shopee/product/entity"
	"github.com/wjp-letgo/letgo/x/api/shopee/shop"
	shopEntity "github.com/wjp-letgo/letgo/x/api/shopee/shop/entity"
)

//Shopeer
type Shopeer interface{
	//auth
	AuthorizationURL()string
	GetAccesstoken(code string,shopID int64) authEntity.GetAccessTokenResult
	RefreshAccessToken(shop commonentity.ShopInfo)authEntity.RefreshAccessTokenResult
	//order
	GetOrderList(
	timeRangeField order.TimeRangeField,
	timeFrom,timeTo,pageSize int,
	cursor string,
	orderStatus order.OrderStatus,
	responseOptionalFields string) orderEntity.GetOrderListResult
	GetShipmentList(cursor string,pageSize int) orderEntity.GetShipmentListResult
	GetOrderDetail(orderSnList []string,responseOptionalFields ...string) orderEntity.GetOrderDetailResult
	SplitOrder(orderSn string,packageList []orderEntity.PackageListRequestEntity) orderEntity.SplitOrderResult
	UnSplitOrder(orderSn string) orderEntity.UnSplitOrderResult
	CancelOrder(orderSn string,cancelReason order.CancelReason,itemList []orderEntity.CancelOrderRequestEntity) orderEntity.CancelOrderResult
	HandleBuyerCancellation(orderSn string,operation order.Operation) orderEntity.HandleBuyerCancellationResult
	SetNote(orderSn,note string) orderEntity.SetNoteResult
	AddInvoiceData(orderSn string,invoiceData orderEntity.InvoiceDataEntity) orderEntity.AddInvoiceDataResult
	//logistics
	GetShippingParameter(orderSn string)logisticsEntity.GetShippingParameterResult
	GetTrackingNumber(orderSn,packageNumber string,responseOptionalFields ...string)logisticsEntity.GetTrackingNumberResult
	ShipOrder(orderSn,packageNumber string,pickup *logisticsEntity.ShipOrderRequestPickupEntity,dropoff *logisticsEntity.ShipOrderRequestDropoffEntity,nonIntegrated *logisticsEntity.ShipOrderRequestNonIntegratedEntity)logisticsEntity.ShipOrderResult
	UpdateShippingOrder(orderSn,packageNumber string,pickup *logisticsEntity.UpdateShippingOrderRequestPickupEntity)logisticsEntity.UpdateShippingOrderResult
	GetShippingDocumentParameter(orderList *logisticsEntity.ShippingDocumentParameterRequestOrderListEntity)logisticsEntity.GetShippingDocumentParameterResult
	CreateShippingDocument(orderList *logisticsEntity.CreateShippingDocumentRequestOrderListEntity)logisticsEntity.CreateShippingDocumentResult
	GetShippingDocumentResult(orderList *logisticsEntity.GetShippingDocumentResultRequestOrderListEntity)logisticsEntity.GetShippingDocumentResult
	DownloadShippingDocument(orderList *logisticsEntity.DownloadShippingDocumentRequestOrderListEntity)logisticsEntity.DownloadShippingDocumentResult
	GetShippingDocumentInfo(orderSn,packageNumber string)logisticsEntity.GetShippingDocumentInfoResult
	GetTrackingInfo(orderSn,packageNumber string)logisticsEntity.GetTrackingInfoResult
	GetAddressList()logisticsEntity.GetAddressListResult
	SetAddressConfig(showPickupAddress bool,AddressTypeConfig logisticsEntity.AddressTypeConfigEntity)logisticsEntity.SetAddressConfigResult
	DeleteAddress(addressID int64)logisticsEntity.DeleteAddressResult
	GetChannelList()logisticsEntity.GetChannelListResult
	UpdateChannel(logisticsChannelID int64,enabled,preferred,codEnabled bool)logisticsEntity.UpdateChannelResult
	BatchShipOrder(orderList *logisticsEntity.BatchShipOrderRequestOrderListEntity,pickup *logisticsEntity.BatchShipOrderRequestPickupEntity,dropoff *logisticsEntity.BatchShipOrderRequestDropoffEntity,nonIntegrated *logisticsEntity.BatchShipOrderRequestNonIntegratedEntity)logisticsEntity.BatchShipOrderResult
	//product
	GetComment(itemID,commentID int64,cursor string,pageSize int)productEntity.GetCommentResult
	ReplyComment(commentList []productEntity.ReplyCommentRequestCommentEntity)productEntity.ReplyCommentResult
	GetItemBaseInfo(itemIdList []int64)productEntity.GetItemBaseInfoResult
	GetItemExtraInfo(itemIdList []int64)productEntity.GetItemExtraInfoResult
	GetItemList(offset,pageSize,updateTimeFrom,updateTimeTo int,itemStatus product.ItemStatus)productEntity.GetItemListResult
	DeleteItem(itemID int64)productEntity.DeleteItemResult
	AddItem(item productEntity.AddItemRequestItemEntity)productEntity.AddItemResult
	UpdateItem(item productEntity.UpdateItemRequestItemEntity)productEntity.UpdateItemResult
	GetModelList(itemID int64)productEntity.GetModelListResult
	UpdateSizeChart(itemID int64,sizeChart string)productEntity.UpdateSizeChartResult
	UnlistItem(itemList []productEntity.UnlistItemItemListEntity)productEntity.UnlistItemResult
	BoostItem(itemIdList []int64)productEntity.BoostItemResult
	GetBoostedList(itemIdList []int64)productEntity.GetBoostedListResult
	GetDtsLimit(categoryID int64)productEntity.GetDtsLimitResult
	GetItemLimit(categoryID int64)productEntity.GetItemLimitResult
	SupportSizeChart(category int64)productEntity.SupportSizeChartResult
	InitTierVariation(itemID int64,tierVariation productEntity.TierVariationEntity,model productEntity.InitTierVariationModelEntity)productEntity.InitTierVariationResult
	UpdateTierVariation(itemID int64,tierVariation []productEntity.TierVariationEntity)productEntity.UpdateTierVariationResult
	UpdateModel(itemID int64,model []productEntity.UpdateModelEntity)productEntity.UpdateModelResult
	AddModel(itemID int64,modelList []productEntity.InitTierVariationModelEntity)productEntity.AddModelResult
	DeleteModel(itemID,modelID int64)productEntity.DeleteModelResult
	UpdatePrice(itemID int64,priceList []productEntity.UpdatePricePriceInfoEntity)productEntity.UpdatePriceResult
	UpdateStock(itemID int64,stockList []productEntity.UpdateStockStockInfoEntity)productEntity.UpdateStockResult
	GetCategory(language string)productEntity.GetCategoryResult
	GetAttributes(language string,categoryID int64)productEntity.GetAttributesResult
	GetBrandList(offset,pageSize int,categoryID int64,status product.BrandStatus)productEntity.GetBrandListResult
	CategoryRecommend(itemName string)productEntity.CategoryRecommendResult
	GetItemPromotion(itemIdList []int64)productEntity.GetItemPromotionResult
	UpdateSipItemPrice(itemID int64,sipItemPrice []productEntity.SipItemPriceEntity)productEntity.UpdateSipItemPriceResult
	SearchItem(offset string,pageSize int,itemName string,attributeStatus product.AttributeStatus)productEntity.SearchItemResult
	//global_product
	GetGlobalCategory(language string)globalProductEntity.GetCategoryResult
	GetGlobalAttributes(language string,categoryID int64)globalProductEntity.GetAttributesResult
	GetGlobalBrandList(offset,pageSize int,categoryID int64,status globalproduct.BrandStatus)globalProductEntity.GetBrandListResult
	GetGlobalItemLimit(categoryID int64)globalProductEntity.GetGlobalItemLimitResult
	GetGlobalDtsLimit(categoryID int64)globalProductEntity.GetDtsLimitResult
	GetGlobalItemList(offset string,pageSize int)globalProductEntity.GetGlobalItemListResult
	GetGlobalItemInfo(globalItemIdList []int64)globalProductEntity.GetGlobalItemInfoResult
	AddGlobalItem(item globalProductEntity.AddItemRequestItemEntity)globalProductEntity.AddGlobalItemResult
	UpdateGlobalItem(item globalProductEntity.UpdateItemRequestItemEntity)globalProductEntity.AddGlobalItemResult
	DeleteGlobalItem(globalItemID int64)globalProductEntity.DeleteGlobalItemResult
	InitGlobalTierVariation(globalItemID int64,tierVariation globalProductEntity.TierVariationEntity,model globalProductEntity.InitTierVariationModelEntity)globalProductEntity.InitTierVariationResult
	UpdateGlobalTierVariation(globalItemID int64,tierVariation []globalProductEntity.TierVariationEntity)globalProductEntity.UpdateTierVariationResult
	AddGlobalModel(globalItemID int64,modelList []globalProductEntity.InitTierVariationModelEntity)globalProductEntity.AddGlobalModelResult
	UpdateGlobalModel(globalItemID int64,model []globalProductEntity.UpdateModelEntity)globalProductEntity.UpdateGlobalModelResult
	DeleteGlobalModel(globalItemID,globalModelID int64)globalProductEntity.DeleteGlobalModelResult
	GetGlobalModelList(globalItemID int64)globalProductEntity.GetModelListResult
	SupportGlobalSizeChart(categoryID int64)globalProductEntity.SupportGlobalSizeChartResult
	UpdateGlobalSizeChart(globalItemID int64,sizeChart string)globalProductEntity.UpdateGlobalSizeChartResult
	CreatePublishTask(globalItemID,shopID int64,shopRegion string,item globalProductEntity.CreatePublishTaskItemEntity)globalProductEntity.CreatePublishTaskResult
	GetPublishableShop(globalItemID int64)globalProductEntity.GetPublishableShopResult
	GetPublishTaskResult(publishTaskID int64)globalProductEntity.GetPublishTaskResult
	GetPublishedList(globalItemID int64)globalProductEntity.GetPublishedListResult
	UpdateGlobalPrice(globalItemID int64,priceList []globalProductEntity.PriceEntity)globalProductEntity.UpdateGlobalPriceResult
	UpdateGlobalStock(globalItemID int64,stockList []globalProductEntity.UpdateStockStockInfoEntity)globalProductEntity.UpdateGlobalStockResult
	SetSyncField(shopSyncList []globalProductEntity.ShopSyncEntity)globalProductEntity.SetSyncFieldResult
	GetGlobalItemID(shopID int64,itemIdList []int64)globalProductEntity.GetGlobalItemIDResult
	//media_space
	InitVideoUpload(fileMd5 string,fileSize int)mediaspaceEntity.InitVideoUploadResult
	UploadVideoPart(videoUploadID string,partSeq int,contentMd5 string,partContentPath string)mediaspaceEntity.UploadVideoPartResult
	CompleteVideoUpload(videoUploadID string,partSeqList []int,reportData mediaspaceEntity.ReportDataEntity)mediaspaceEntity.CompleteVideoUploadResult
	GetVideoUploadResult(videoUploadID string)mediaspaceEntity.GetVideoUploadResult
	CancelVideoUpload(videoUploadID string)mediaspaceEntity.CancelVideoUploadResult
	UploadImage(image string)mediaspaceEntity.UploadImageResult
	//shop
	GetShopInfo()shopEntity.GetShopInfoResult
	GetProfile()shopEntity.GetProfileResult
	UpdateProfile(shopName,shopLogo,description string)shopEntity.UpdateProfileResult
	//merchant
	GetMerchantInfo()merchantEntity.GetMerchantInfoResult
	GetShopListByMerchant(pageNo,pageSize int)merchantEntity.GetShopListByMerchantResult
}
//Shopee
type Shopee struct{
	auth.Auth
	order.Order
	logistics.Logistics
	product.Product
	globalproduct.GlobalProduct
	mediaspace.MediaSpace
	shop.Shop
	merchant.Merchant
}

//shopeeList 接口列表
var shopeeList map[string]Shopeer

//Register
func Register(name string,cfg *shopeeConfig.Config){
	shopeeList[name]=&Shopee{
		auth.Auth{Config:cfg},
		order.Order{Config:cfg},
		logistics.Logistics{Config:cfg},
		product.Product{Config:cfg},
		globalproduct.GlobalProduct{Config:cfg},
		mediaspace.MediaSpace{Config:cfg},
		shop.Shop{Config:cfg},
		merchant.Merchant{Config:cfg},
	}
}
//GetApi
func GetApi(name string)Shopeer{
	return shopeeList[name];
}

//init
func init(){
	shopeeList=make(map[string]Shopeer)
}