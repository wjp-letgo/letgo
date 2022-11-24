package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//PreOrderEntity
type PreOrderEntity struct{
	DaysToShip int `json:"days_to_ship"`
}

//String
func(p PreOrderEntity)String()string{
	return lib.ObjectToString(p)
}

type CreatePublishTaskPreOrderEntity struct{
	DaysToShip int `json:"days_to_ship"`
	IsPreOrder bool `json:"is_pre_order"`
}

//String
func(p CreatePublishTaskPreOrderEntity)String()string{
	return lib.ObjectToString(p)
}