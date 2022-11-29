package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemEntity
type ItemEntity struct {
	ItemID        int64             `json:"item_id"`
	CategoryID    int64             `json:"category_id"`
	ItemName      string            `json:"item_name"`
	Description   string            `json:"description"`
	ItemSku       string            `json:"item_sku"`
	CreateTime    int               `json:"create_time"`
	UpdateTime    int               `json:"update_time"`
	AttributeList []AttributeEntity `json:"attribute_list"`
	PriceInfo     []PriceInfoEntity `json:"price_info"`
	StockInfo     []StockInfoEntity `json:"stock_info"`
	Image         ImageEntity       `json:"image"`
	Weight        string            `json:"weight"`
	Dimension     DimensionEntity   `json:"dimension"`
	PreOrder      PreOrderEntity    `json:"pre_order"`
	Condition     string            `json:"condition"`
	SizeChart     string            `json:"size_chart"`
	ItemStatus    string            `json:"item_status"`
	HasModel      bool              `json:"has_model"`
	PromotionID   int64             `json:"promotion_id"`
	Brand         BrandEntity       `json:"brand"`
	ItemDangerous int               `json:"item_dangerous"`
}

//String
func (i ItemEntity) String() string {
	return lib.ObjectToString(i)
}

//ItemEntity
type ItemExtraEntity struct {
	ItemID       int64   `json:"item_id"`
	Sale         int     `json:"sale"`
	Views        int     `json:"views"`
	Likes        int     `json:"likes"`
	RatingStar   float32 `json:"rating_star"`
	CommentCount int     `json:"comment_count"`
}

//String
func (i ItemExtraEntity) String() string {
	return lib.ObjectToString(i)
}

//ItemListEntity
type ItemListEntity struct {
	ItemID     int64  `json:"item_id"`
	ItemStatus string `json:"item_status"`
	UpdateTime int    `json:"update_time"`
}

//String
func (i ItemListEntity) String() string {
	return lib.ObjectToString(i)
}

//AddItemRequestItemEntity
type AddItemRequestItemEntity struct {
	OriginalPrice  float32           `json:"original_price "`
	Description    string            `json:"description"`
	Weight         float32           `json:"weight"`
	GlobalItemName string            `json:"global_item_name"`
	Dimension      DimensionEntity   `json:"dimension"`
	NormalStock    int               `json:"normal_stock"`
	AttributeList  []AttributeEntity `json:"attribute_list"`
	CategoryID     int64             `json:"category_id"`
	Image          ImageEntity       `json:"image"`
	PreOrder       PreOrderEntity    `json:"pre_order"`
	GlobalItemSku  string            `json:"global_item_sku"`
	Condition      string            `json:"condition"`
	VideoUploadID  []string          `json:"video_upload_id"`
	Brand          BrandEntity       `json:"brand"`
}

//String
func (i AddItemRequestItemEntity) String() string {
	return lib.ObjectToString(i)
}

//UnlistItemItemListEntity
type UnlistItemItemListEntity struct {
	ItemID int64 `json:"item_id"`
	Unlist bool  `json:"unlist"`
}

//String
func (i UnlistItemItemListEntity) String() string {
	return lib.ObjectToString(i)
}

//GetBoostedListItemListEntity
type GetBoostedListItemListEntity struct {
	ItemID         int64 `json:"item_id"`
	CoolDownSecond int   `json:"cool_down_second"`
}

//String
func (i GetBoostedListItemListEntity) String() string {
	return lib.ObjectToString(i)
}

//UpdateItemRequestItemEntity
type UpdateItemRequestItemEntity struct {
	GlobalItemID   int64             `json:"global_item_id"`
	OriginalPrice  float32           `json:"original_price "`
	Description    string            `json:"description"`
	Weight         float32           `json:"weight"`
	GlobalItemName string            `json:"global_item_name"`
	Dimension      DimensionEntity   `json:"dimension"`
	NormalStock    int               `json:"normal_stock"`
	AttributeList  []AttributeEntity `json:"attribute_list"`
	CategoryID     int64             `json:"category_id"`
	Image          ImageEntity       `json:"image"`
	PreOrder       PreOrderEntity    `json:"pre_order"`
	GlobalItemSku  string            `json:"global_item_sku"`
	Condition      string            `json:"condition"`
	VideoUploadID  []string          `json:"video_upload_id"`
	Brand          BrandEntity       `json:"brand"`
}

//String
func (i UpdateItemRequestItemEntity) String() string {
	return lib.ObjectToString(i)
}

//CreatePublishTaskItemEntity
type CreatePublishTaskItemEntity struct {
	item_name      string                          `json:"item_name"`
	description    string                          `json:"description"`
	item_status    string                          `json:"item_status"`
	original_price float32                         `json:"original_price"`
	image          ImageEntity                     `json:"image"`
	category_id    int64                           `json:"category_id"`
	tier_variation []TierVariationEntity           `json:"tier_variation"`
	model          []CreatePublishTaskModelEntity  `json:"model"`
	size_chart     string                          `json:"size_chart"`
	logistic       []LogisticEntity                `json:"logistic"`
	PreOrder       CreatePublishTaskPreOrderEntity `json:"pre_order"`
}

//String
func (i CreatePublishTaskItemEntity) String() string {
	return lib.ObjectToString(i)
}
