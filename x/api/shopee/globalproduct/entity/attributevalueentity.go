package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//AttributeValueEntity
type AttributeValueEntity struct {
	ValueID             int64                   `json:"value_id"`
	OriginalValueName   string                  `json:"original_value_name"`
	ValueUnit           string                  `json:"value_unit"`
	DisplayValueName    string                  `json:"display_value_name"`
	ParentAttributeList []ParentAttributeEntity `json:"parent_attribute_list"`
	ParentBrandList     []ParentBrandEntity     `json:"parent_brand_list"`
}

//String
func (a AttributeValueEntity) String() string {
	return lib.ObjectToString(a)
}
