package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//CategoryEntity
type CategoryEntity struct {
	CategoryID           int64  `json:"category_id"`
	ParentCategoryID     int64  `json:"parent_category_id"`
	OriginalCategoryName string `json:"original_category_name"`
	DisplayCategoryName  string `json:"display_category_name"`
	HasChildren          bool   `json:"has_children"`
}

//String
func (c CategoryEntity) String() string {
	return lib.ObjectToString(c)
}
