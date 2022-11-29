package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//AttributeEntity
type AttributeEntity struct {
	AttributeID           int64                  `json:"attribute_id"`
	OriginalAttributeName string                 `json:"original_attribute_name"`
	DisplayAttributeName  string                 `json:"display_attribute_name"`
	InputValidationType   string                 `json:"input_validation_type"`
	FormatType            string                 `json:"format_type"`
	DateFormatType        string                 `json:"date_format_type"`
	InputType             string                 `json:"input_type"`
	AttributeUnit         []string               `json:"attribute_unit"`
	IsMandatory           bool                   `json:"is_mandatory"`
	AttributeType         int                    `json:"attribute_type"`
	AttributeValueList    []AttributeValueEntity `json:"attribute_value_list"`
}

//String
func (a AttributeEntity) String() string {
	return lib.ObjectToString(a)
}
