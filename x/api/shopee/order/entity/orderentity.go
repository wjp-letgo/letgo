package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"strings"
	"reflect"
)

//OrderEntity
type OrderEntity struct{
	OrderSN string `json:"order_sn"`
	Region string `json:"region"`
	Currency string `json:"currency"`
	Cod bool `json:"cod"`
	TotalAmount float32 `json:"total_amount"`
	OrderStatus string `json:"order_status"`
	ShippingCarrier string `json:"shipping_carrier"`
	PaymentMethod string `json:"payment_method"`
	EstimatedShippingFee float32 `json:"estimated_shipping_fee"`
	MessageToSeller string `json:"message_to_seller"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
	DaysToShip int `json:"days_to_ship"`
	ShipByDate int `json:"ship_by_date"`
	BuyerUserID int `json:"buyer_user_id"`
	BuyerUsername string `json:"buyer_username"`
	RecipientAddress RecipientAddressEntity `json:"recipient_address"`
	ActualShippingFee float32 `json:"actual_shipping_fee"`
	GoodsToDeclare bool `json:"goods_to_declare"`
	Note string `json:"note"`
	NoteUpdateTime int `json:"note_update_time"`
	ItemList []ItemListEntity `json:"item_list"`
	PayTime int `json:"pay_time"`
	Dropshipper string `json:"dropshipper"`
	CreditCardNumber string `json:"credit_card_number"`
	DropshipperPhone string `json:"dropshipper_phone"`
	SplitUp bool `json:"split_up"`
	BuyerCancelReason string `json:"buyer_cancel_reason"`
	CancelBy string `json:"cancel_by"`
	CancelReason string `json:"cancel_reason"`
	ActualShippingFeeConfirmed bool `json:"actual_shipping_fee_confirmed"`
	BuyerCpfID string `json:"buyer_cpf_id"`
	FulfillmentFlag string `json:"fulfillment_flag"`
	PickupDoneTime int `json:"pickup_done_time"`
	PackageList []PackageListEntity `json:"package_list"`
	InvoiceData InvoiceDataEntity `json:"invoice_data"`
	CheckoutShippingCarrier string `json:"checkout_shipping_carrier"`
}

//String
func(o OrderEntity)String()string{
	return lib.ObjectToString(o)
}

//ResponseOptionalFields
func OrderResponseOptionalFields()string{
	var fields []string
	order:=OrderEntity{}
	orderType:=reflect.TypeOf(order)
	for i:=0;i<orderType.NumField();i++{
		fields=append(fields, orderType.Field(i).Tag.Get("json"))
	}
	if len(fields)>0{
		return strings.Join(fields,",")
	}
	return ""
}