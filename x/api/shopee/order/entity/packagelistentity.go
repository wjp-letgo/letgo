package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PackageListEntity
type PackageListEntity struct {
	PackageNumber   string           `json:"package_number"`
	LogisticsStatus string           `json:"logistics_status"`
	ShippingCarrier string           `json:"shipping_carrier"`
	ItemList        []ItemListEntity `json:"item_list"`
}

//String
func (p PackageListEntity) String() string {
	return lib.ObjectToString(p)
}

//PackageListRequestEntity
type PackageListRequestEntity struct {
	ItemList []PackageListRequestItemListEntity `json:"item_list"`
}

//String
func (p PackageListRequestEntity) String() string {
	return lib.ObjectToString(p)
}
