package logistics

import (
	"strings"

	"github.com/wjp-letgo/letgo/lib"
	shopeeConfig "github.com/wjp-letgo/letgo/x/api/shopee/config"
	"github.com/wjp-letgo/letgo/x/api/shopee/logistics/entity"
)

//常量
const (
	NORMAL_AIR_WAYBILL ShippingDocumentType="NORMAL_AIR_WAYBILL"
	THERMAL_AIR_WAYBILL ShippingDocumentType="THERMAL_AIR_WAYBILL"
	NORMAL_JOB_AIR_WAYBILL ShippingDocumentType="NORMAL_JOB_AIR_WAYBILL"
	THERMAL_JOB_AIR_WAYBILL ShippingDocumentType="THERMAL_JOB_AIR_WAYBILL"
	READY ShippingDocumentStatus ="READY"
	FAILED ShippingDocumentStatus ="FAILED"
	PROCESSING ShippingDocumentStatus ="PROCESSING"
)
//ShippingDocumentType
type ShippingDocumentType string

//ShippingDocumentStatus
type ShippingDocumentStatus string
//Logistics
type Logistics struct{
	Config *shopeeConfig.Config
}

//GetShippingParameter
//@Title Use this api to get shipping parameter.
//@Description https://open.shopee.com/documents?module=95&type=1&id=550&version=2
func (l *Logistics)GetShippingParameter(orderSn string)entity.GetShippingParameterResult{
	method:="logistics/get_shipping_parameter"
	params:=lib.InRow{
		"order_sn":orderSn,
	}
	result:=entity.GetShippingParameterResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetTrackingNumber
//@Title Use this api to get tracking_number when you hava shipped order.
//@Description https://open.shopee.com/documents?module=95&type=1&id=552&version=2
func (l *Logistics)GetTrackingNumber(orderSn,packageNumber string,responseOptionalFields ...string)entity.GetTrackingNumberResult{
	method:="logistics/get_tracking_number"
	responseOptionalFieldsStr:=""
	if len(responseOptionalFields)>0{
		responseOptionalFieldsStr=strings.Join(responseOptionalFields,",")
	}else{
		responseOptionalFieldsStr=entity.TrackingNumberResponseOptionalFields()
	}
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_number":packageNumber,
		"response_optional_fields":responseOptionalFieldsStr,
	}
	result:=entity.GetTrackingNumberResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//ShipOrder
//@Title Use this api to initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller.
//@Description https://open.shopee.com/documents?module=95&type=1&id=553&version=2
func (l *Logistics)ShipOrder(orderSn,packageNumber string,pickup *entity.ShipOrderRequestPickupEntity,dropoff *entity.ShipOrderRequestDropoffEntity,nonIntegrated *entity.ShipOrderRequestNonIntegratedEntity)entity.ShipOrderResult{
	method:="logistics/ship_order"
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_number":packageNumber,
	}
	if pickup!=nil{
		params["pickup"]=*pickup
	}
	if dropoff!=nil{
		params["dropoff"]=*dropoff
	}
	if nonIntegrated!=nil{
		params["non_integrated"]=*nonIntegrated
	}
	result:=entity.ShipOrderResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//UpdateShippingOrder
//@Title Use this api to update pickup address and pickup time for shipping order.
//@Description https://open.shopee.com/documents?module=95&type=1&id=555&version=2
func (l *Logistics)UpdateShippingOrder(orderSn,packageNumber string,pickup *entity.UpdateShippingOrderRequestPickupEntity)entity.UpdateShippingOrderResult{
	method:="logistics/update_shipping_order"
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_number":packageNumber,
	}
	if pickup!=nil{
		params["pickup"]=*pickup
	}
	result:=entity.UpdateShippingOrderResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetShippingDocumentParameter
//@Title Use this api to get the selectable shipping_document_type and suggested shipping_document_type.
//@Description https://open.shopee.com/documents?module=95&type=1&id=549&version=2
func (l *Logistics)GetShippingDocumentParameter(orderList *entity.ShippingDocumentParameterRequestOrderListEntity)entity.GetShippingDocumentParameterResult{
	method:="logistics/get_shipping_document_parameter"
	params:=lib.InRow{
		"order_list":*orderList,
	}
	result:=entity.GetShippingDocumentParameterResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//CreateShippingDocument
//@Title Use this api to create shipping document task for each order or package
//@Description https://open.shopee.com/documents?module=95&type=1&id=547&version=2
func (l *Logistics)CreateShippingDocument(orderList *entity.CreateShippingDocumentRequestOrderListEntity)entity.CreateShippingDocumentResult{
	method:="logistics/create_shipping_document"
	params:=lib.InRow{
		"order_list":*orderList,
	}
	result:=entity.CreateShippingDocumentResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetShippingDocumentResult
//@Title Use this api to get the shipping_document of each order or package status.
//@Description https://open.shopee.com/documents?module=95&type=1&id=561&version=2
func (l *Logistics)GetShippingDocumentResult(orderList *entity.GetShippingDocumentResultRequestOrderListEntity)entity.GetShippingDocumentResult{
	method:="logistics/get_shipping_document_result"
	params:=lib.InRow{
		"order_list":*orderList,
	}
	result:=entity.GetShippingDocumentResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DownloadShippingDocument
//@Title Use this api to download shipping_document. You have to call v2.logistics.create_shipping_document to create a new shipping document task first and call v2.logistics.get_shipping_document_resut to get the task status second. If the task is READY, you can download this shipping document.
//@Description https://open.shopee.com/documents?module=95&type=1&id=548&version=2
func (l *Logistics)DownloadShippingDocument(orderList *entity.DownloadShippingDocumentRequestOrderListEntity)entity.DownloadShippingDocumentResult{
	method:="logistics/download_shipping_document"
	params:=lib.InRow{
		"order_list":*orderList,
	}
	result:=entity.DownloadShippingDocumentResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetShippingDocumentInfo
//@Title Use this api to fetch the logistics information of an order, these info can be used for airwaybill printing. Dedicated for crossborder SLS order airwaybill. May not be applicable for local channel airwaybill.
//@Description https://open.shopee.com/documents?module=95&type=1&id=560&version=2
func (l *Logistics)GetShippingDocumentInfo(orderSn,packageNumber string)entity.GetShippingDocumentInfoResult{
	method:="logistics/get_shipping_document_info"
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_number":packageNumber,
	}
	result:=entity.GetShippingDocumentInfoResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetTrackingInfo
//@Title Use this api to get the logistics tracking information of an order.
//@Description https://open.shopee.com/documents?module=95&type=1&id=551&version=2
func (l *Logistics)GetTrackingInfo(orderSn,packageNumber string)entity.GetTrackingInfoResult{
	method:="logistics/get_tracking_info"
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_number":packageNumber,
	}
	result:=entity.GetTrackingInfoResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//GetAddressList
//@Title For integrated logistics channel, use this call to get pickup address for pickup mode order.
//@Description https://open.shopee.com/documents?module=95&type=1&id=558&version=2
func (l *Logistics)GetAddressList()entity.GetAddressListResult{
	method:="logistics/get_address_list"
	params:=lib.InRow{
	}
	result:=entity.GetAddressListResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//SetAddressConfig
//@Title Use this API to set address config of your shop.
//@Description https://open.shopee.com/documents?module=95&type=1&id=556&version=2
func (l *Logistics)SetAddressConfig(showPickupAddress bool,AddressTypeConfig entity.AddressTypeConfigEntity)entity.SetAddressConfigResult{
	method:="logistics/set_address_config"
	params:=lib.InRow{
		"show_pickup_address":showPickupAddress,
		"address_type_config":AddressTypeConfig,
	}
	result:=entity.SetAddressConfigResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//DeleteAddress
//@Title Use this api to delete address.
//@Description https://open.shopee.com/documents?module=95&type=1&id=598&version=2
func (l *Logistics)DeleteAddress(addressID int64)entity.DeleteAddressResult{
	method:="logistics/delete_address"
	params:=lib.InRow{
		"address_id":addressID,
	}
	result:=entity.DeleteAddressResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetChannelList
//@Title Use this api to get all supported logistic channels.
//@Description https://open.shopee.com/documents?module=95&type=1&id=559&version=2
func (l *Logistics)GetChannelList()entity.GetChannelListResult{
	method:="logistics/get_channel_list"
	params:=lib.InRow{
	}
	result:=entity.GetChannelListResult{}
	err:=l.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UpdateChannel
//@Title Use this api to update shop level logistics channel's configuration.
//@Description https://open.shopee.com/documents?module=95&type=1&id=554&version=2
func (l *Logistics)UpdateChannel(logisticsChannelID int64,enabled,preferred,codEnabled bool)entity.UpdateChannelResult{
	method:="logistics/update_channel"
	params:=lib.InRow{
		"logistics_channel_id":logisticsChannelID,
		"enabled":enabled,
		"preferred":preferred,
		"cod_enabled":codEnabled,
	}
	result:=entity.UpdateChannelResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//BatchShipOrder
//@Title Use this api to batch initiate logistics including arrange pickup, dropoff or shipment for non-integrated logistic channels. Should call v2.logistics.get_shipping_parameter to fetch all required param first. It's recommended to initiate logistics one hour after the orders were placed since there is one-hour window buyer can cancel any order without request to seller.
//@Description https://open.shopee.com/documents?module=95&type=1&id=688&version=2
func (l *Logistics)BatchShipOrder(orderList *entity.BatchShipOrderRequestOrderListEntity,pickup *entity.BatchShipOrderRequestPickupEntity,dropoff *entity.BatchShipOrderRequestDropoffEntity,nonIntegrated *entity.BatchShipOrderRequestNonIntegratedEntity)entity.BatchShipOrderResult{
	method:="logistics/batch_ship_order"
	params:=lib.InRow{
		"order_list":*orderList,
	}
	if pickup!=nil{
		params["pickup"]=*pickup
	}
	if dropoff!=nil{
		params["dropoff"]=*dropoff
	}
	if nonIntegrated!=nil{
		params["non_integrated"]=*nonIntegrated
	}
	result:=entity.BatchShipOrderResult{}
	err:=l.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}