package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemNameLengthLimitEntity
type ItemNameLengthLimitEntity struct {
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func (p ItemNameLengthLimitEntity) String() string {
	return lib.ObjectToString(p)
}
