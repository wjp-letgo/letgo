package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//DropoffEntity
type DropoffEntity struct{
	BranchList []BranchEntity `json:"branch_list"`
}

//String
func(d DropoffEntity)String()string{
	return lib.ObjectToString(d)
}


//ShipOrderRequestDropoffEntity
type ShipOrderRequestDropoffEntity struct{
	BranchID int64 `json:"branch_id"`
	SenderRealName string `json:"sender_real_name"`
	TrackingNumber string `json:"tracking_number"`
}

//String
func(s ShipOrderRequestDropoffEntity)String()string{
	return lib.ObjectToString(s)
}


//BatchShipOrderRequestDropoffEntity
type BatchShipOrderRequestDropoffEntity struct{
	BranchID int64 `json:"branch_id"`
	SenderRealName string `json:"sender_real_name"`
	TrackingNumber string `json:"tracking_number"`
}

//String
func(b BatchShipOrderRequestDropoffEntity)String()string{
	return lib.ObjectToString(b)
}