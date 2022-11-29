package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ParentAttributeEntity
type ParentAttributeEntity struct {
	ParentAttributeID int64 `json:"parent_attribute_id"`
	ParentValueID     int64 `json:"parent_value_id"`
}

//String
func (a ParentAttributeEntity) String() string {
	return lib.ObjectToString(a)
}
