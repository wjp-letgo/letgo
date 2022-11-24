package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ThirdPartyLogisticInfoEntity
type ThirdPartyLogisticInfoEntity struct{
	ServiceDescription string `json:"service_description"`
	Barcode string `json:"barcode"`
	PurchaseTime string `json:"purchase_time"`
	ReturnTime string `json:"return_time"`
	ManufacturersName string `json:"manufacturers_name"`
	ManufacturersWebsite string `json:"manufacturers_website"`
	RecipientArea string `json:"recipient_area"`
	RouteStep string `json:"route_step"`
	Suda5Code string `json:"suda5_code"`
	LargeLogisticsID string `json:"large_logistics_id"`
	ParentID string `json:"parent_id"`
	ReturnCycle string `json:"return_cycle"`
	ReturnMode string `json:"return_mode"`
	Prompt string `json:"prompt"`
	OrderSn string `json:"order_sn"`
	Qrcode string `json:"qrcode"`
	EcSupplierName string `json:"ec_supplier_name"`
	EcBarCode16 string `json:"ec_bar_code16"`
	EquipmentID string `json:"equipment_id"`
	EshopID string `json:"eshop_id"`
	EcBarCode9 string `json:"ec_bar_code9"`
	PelicanTrackingNo string `json:"pelican_tracking_no"`
	PrintDate string `json:"print_date"`
	Pzip string `json:"pzip"`
	PzipC string `json:"pzip_c"`
	DeliverAreaTxt string `json:"deliver_area_txt"`
	DeliverDateYmd string `json:"deliver_date_ymd"`
	SdDriverCode string `json:"sd_driver_code"`
	MdDriverCode string `json:"md_driver_code"`
	PutorderStackzoneCode string `json:"putorder_stackzone_code"`
	CustomerCode string `json:"customer_code"`
}

//String
func(t ThirdPartyLogisticInfoEntity)String()string{
	return lib.ObjectToString(t)
}
