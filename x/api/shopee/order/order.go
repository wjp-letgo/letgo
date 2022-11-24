package order

import (
	shopeeConfig "github.com/wjpxxx/letgo/x/api/shopee/config"
	"github.com/wjpxxx/letgo/x/api/shopee/order/entity"
	"github.com/wjpxxx/letgo/lib"
	"strings"
)
const (
	CREATE_TIME TimeRangeField="create_time"
	UPDATE_TIME TimeRangeField="update_time"
	UNPAID OrderStatus="UNPAID"
	READY_TO_SHIP OrderStatus="READY_TO_SHIP"
	PROCESSED OrderStatus="PROCESSED"
	RETRY_SHIP OrderStatus="RETRY_SHIP"
	SHIPPED OrderStatus="SHIPPED"
	TO_CONFIRM_RECEIVE OrderStatus="TO_CONFIRM_RECEIVE"
	COMPLETED OrderStatus="COMPLETED"
	IN_CANCEL OrderStatus="IN_CANCEL"
	CANCELLED OrderStatus="CANCELLED"
	TO_RETURN OrderStatus="TO_RETURN"
	OUT_OF_STOCK CancelReason = "OUT_OF_STOCK"
	CUSTOMER_REQUEST CancelReason = "CUSTOMER_REQUEST"
	UNDELIVERABLE_AREA CancelReason = "UNDELIVERABLE_AREA"
	COD_NOT_SUPPORTED CancelReason = "COD_NOT_SUPPORTED"
	ACCEPT Operation = "ACCEPT"
	REJECT Operation = "REJECT"
)
//TimeRangeField
type TimeRangeField string

//OrderStatus
type OrderStatus string

//CancelReason
type CancelReason string

//Operation
type Operation string

//Order
type Order struct{
	Config *shopeeConfig.Config
}

//GetOrderList
//@Title Use this api to search orders.
//@Description https://open.shopee.com/documents?module=94&type=1&id=542&version=2
func (o *Order)GetOrderList(
	timeRangeField TimeRangeField,
	timeFrom,timeTo,pageSize int,
	cursor string,
	orderStatus OrderStatus,
	responseOptionalFields string) entity.GetOrderListResult{
	method:="order/get_order_list"
	params:=lib.InRow{
		"time_range_field":timeRangeField,
		"time_from":timeFrom,
		"time_to":timeTo,
		"page_size":pageSize,
		"cursor":cursor,
		"order_status":orderStatus,
		"response_optional_fields":responseOptionalFields,
	}
	result:=entity.GetOrderListResult{}
	err:=o.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetShipmentList
//@Title Use this api to get order list which order_status is READY_TO_SHIP.
//@Description https://open.shopee.com/documents?module=94&type=1&id=543&version=2
func (o *Order)GetShipmentList(cursor string,pageSize int) entity.GetShipmentListResult{
	method:="order/get_shipment_list"
	params:=lib.InRow{
		"cursor":cursor,
		"page_size":pageSize,
	}
	result:=entity.GetShipmentListResult{}
	err:=o.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//GetOrderDetail
//@Title Use this api to get order detail.
//@Description https://open.shopee.com/documents?module=94&type=1&id=557&version=2
func (o *Order)GetOrderDetail(orderSnList []string,responseOptionalFields ...string) entity.GetOrderDetailResult{
	method:="order/get_order_detail"
	responseOptionalFieldsStr:=""
	if len(responseOptionalFields)>0{
		responseOptionalFieldsStr=strings.Join(responseOptionalFields,",")
	}else{
		responseOptionalFieldsStr=entity.OrderResponseOptionalFields()
	}
	params:=lib.InRow{
		"order_sn_list":strings.Join(orderSnList,","),
		"response_optional_fields":responseOptionalFieldsStr,
	}
	result:=entity.GetOrderDetailResult{}
	err:=o.Config.HttpGet(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//SplitOrder
//@Title Use this api to split an order into multiple packages.
//@Description https://open.shopee.com/documents?module=94&type=1&id=545&version=2
func (o *Order)SplitOrder(orderSn string,packageList []entity.PackageListRequestEntity) entity.SplitOrderResult{
	method:="order/split_order"
	params:=lib.InRow{
		"order_sn":orderSn,
		"package_list":packageList,
	}
	result:=entity.SplitOrderResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//UnSplitOrder
//@Title Use this ai to undo split of order. After undo split, the order will have only one package.
//@Description https://open.shopee.com/documents?module=94&type=1&id=546&version=2
func (o *Order)UnSplitOrder(orderSn string) entity.UnSplitOrderResult{
	method:="order/unsplit_order"
	params:=lib.InRow{
		"order_sn":orderSn,
	}
	result:=entity.UnSplitOrderResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//CancelOrder
//@Title Use this api to cancel an order.
//@Description https://open.shopee.com/documents?module=94&type=1&id=541&version=2
func (o *Order)CancelOrder(orderSn string,cancelReason CancelReason,itemList []entity.CancelOrderRequestEntity) entity.CancelOrderResult{
	method:="order/cancel_order"
	params:=lib.InRow{
		"order_sn":orderSn,
		"cancel_reason":cancelReason,
		"item_list":itemList,
	}
	result:=entity.CancelOrderResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}


//HandleBuyerCancellation
//@Title Use this api to handle buyer's cancellation application.
//@Description https://open.shopee.com/documents?module=94&type=1&id=544&version=2
func (o *Order)HandleBuyerCancellation(orderSn string,operation Operation) entity.HandleBuyerCancellationResult{
	method:="order/handle_buyer_cancellation"
	params:=lib.InRow{
		"order_sn":orderSn,
		"operation":operation,
	}
	result:=entity.HandleBuyerCancellationResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//SetNote
//@Title Use this api to set note for an order.
//@Description https://open.shopee.com/documents?module=94&type=1&id=540&version=2
func (o *Order)SetNote(orderSn,note string) entity.SetNoteResult{
	method:="order/set_note"
	params:=lib.InRow{
		"order_sn":orderSn,
		"note":note,
	}
	result:=entity.SetNoteResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}

//AddInvoiceData
//@Title Use the API to add invoice data of the order when the order status is PENDING_INVOICE under some special logistics channels, such as Total Express, only for the BR local seller.
//@Description https://open.shopee.com/documents?module=94&type=1&id=685&version=2
func (o *Order)AddInvoiceData(orderSn string,invoiceData entity.InvoiceDataEntity) entity.AddInvoiceDataResult{
	method:="order/set_note"
	params:=lib.InRow{
		"order_sn":orderSn,
		"invoice_data":invoiceData,
	}
	result:=entity.AddInvoiceDataResult{}
	err:=o.Config.HttpPost(method,params,&result)
	if err!=nil{
		result.Error=err.Error()
	}
	return result
}