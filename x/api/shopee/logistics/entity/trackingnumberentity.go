package entity

import (
	"github.com/wjpxxx/letgo/lib"
	"strings"
	"reflect"
)

//TrackingNumberEntity
type TrackingNumberEntity struct{
	TrackingNumber string `json:"tracking_number"`
	PlpNumber string `json:"plp_number"`
	FirstMileTrackingNumber string `json:"first_mile_tracking_number"`
	LastMileTrackingNumber string `json:"last_mile_tracking_number"`
}

//String
func(t TrackingNumberEntity)String()string{
	return lib.ObjectToString(t)
}

//TrackingNumberResponseOptionalFields
func TrackingNumberResponseOptionalFields()string{
	var fields []string
	enty:=TrackingNumberEntity{}
	enType:=reflect.TypeOf(enty)
	for i:=0;i<enType.NumField();i++{
		fields=append(fields, enType.Field(i).Tag.Get("json"))
	}
	if len(fields)>0{
		return strings.Join(fields,",")
	}
	return ""
}