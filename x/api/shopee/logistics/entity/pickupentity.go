package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PickupEntity
type PickupEntity struct{
	AddressList []AddressEntity `json:"address_list"`
}

//String
func(p PickupEntity)String()string{
	return lib.ObjectToString(p)
}

//ShipOrderRequestPickupEntity
type ShipOrderRequestPickupEntity struct{
	AddressID int64 `json:"address_id"`
	PickupTimeID string `json:"pickup_time_id"`
	TrackingNumber string `json:"tracking_number"`
}

//String
func(s ShipOrderRequestPickupEntity)String()string{
	return lib.ObjectToString(s)
}


//UpdateShippingOrderRequestPickupEntity
type UpdateShippingOrderRequestPickupEntity struct{
	AddressID int64 `json:"address_id"`
	PickupTimeID string `json:"pickup_time_id"`
}

//String
func(u UpdateShippingOrderRequestPickupEntity)String()string{
	return lib.ObjectToString(u)
}


//BatchShipOrderRequestPickupEntity
type BatchShipOrderRequestPickupEntity struct{
	AddressID int64 `json:"address_id"`
	PickupTimeID string `json:"pickup_time_id"`
	TrackingNumber string `json:"tracking_number"`
}

//String
func(b BatchShipOrderRequestPickupEntity)String()string{
	return lib.ObjectToString(b)
}
