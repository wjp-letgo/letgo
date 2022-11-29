package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//PriceLimitEntity
type PriceLimitEntity struct {
	MinLimit float32 `json:"min_limit"`
	MaxLimit float32 `json:"max_limit"`
}

//String
func (p PriceLimitEntity) String() string {
	return lib.ObjectToString(p)
}
