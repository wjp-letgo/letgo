package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PreOrderEntity
type PreOrderEntity struct {
	IsPreOrder bool `json:"is_pre_order"`
	DaysToShip int  `json:"days_to_ship"`
}

//String
func (p PreOrderEntity) String() string {
	return lib.ObjectToString(p)
}
