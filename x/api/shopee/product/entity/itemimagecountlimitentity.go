package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemImageCountLimitEntity
type ItemImageCountLimitEntity struct {
	MinLimit int `json:"min_limit"`
	MaxLimit int `json:"max_limit"`
}

//String
func (p ItemImageCountLimitEntity) String() string {
	return lib.ObjectToString(p)
}
