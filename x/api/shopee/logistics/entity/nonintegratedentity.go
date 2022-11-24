package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//NonIntegratedEntity
type ShipOrderRequestNonIntegratedEntity struct{
	TrackingNumber string `json:"tracking_number"`
}

//String
func(n ShipOrderRequestNonIntegratedEntity)String()string{
	return lib.ObjectToString(n)
}


//BatchShipOrderRequestNonIntegratedEntity
type BatchShipOrderRequestNonIntegratedEntity struct{
	TrackingNumber string `json:"tracking_number"`
}

//String
func(b BatchShipOrderRequestNonIntegratedEntity)String()string{
	return lib.ObjectToString(b)
}