package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ShippingDocumentInfoEntity
type ShippingDocumentInfoEntity struct {
	LogisticsChannelID     int64                        `json:"logistics_channel_id"`
	ShippingCarrier        string                       `json:"shipping_carrier"`
	ServiceCode            string                       `json:"service_code"`
	FirstMileName          string                       `json:"first_mile_name"`
	LastMileName           string                       `json:"last_mile_name"`
	GoodsToDeclare         bool                         `json:"goods_to_declare"`
	TrackingNumber         string                       `json:"tracking_number"`
	Zone                   string                       `json:"zone"`
	LaneCode               string                       `json:"lane_code"`
	WarehouseAddress       string                       `json:"warehouse_address"`
	WarehouseID            string                       `json:"warehouse_id"`
	RecipientAddress       RecipientAddressEntity       `json:"recipient_address"`
	Cod                    bool                         `json:"cod"`
	RecipientSortCode      RecipientSortCodeEntity      `json:"recipient_sort_code"`
	SenderSortCode         SenderSortCodeEntity         `json:"sender_sort_code"`
	ThirdPartyLogisticInfo ThirdPartyLogisticInfoEntity `json:"third_party_logistic_info"`
	BuyerCpfID             string                       `json:"buyer_cpf_id"`
	ShopeeTrackingNumber   string                       `json:"shopee_tracking_number"`
	LastMileTrackingNumber string                       `json:"last_mile_tracking_number"`
}

//String
func (s ShippingDocumentInfoEntity) String() string {
	return lib.ObjectToString(s)
}
