package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//WeightLimitEntity
type WeightLimitEntity struct {
	ItemMaxWeight float32 `json:"item_max_weight"`
	ItemMinWeight float32 `json:"item_min_weight"`
}

//String
func (w WeightLimitEntity) String() string {
	return lib.ObjectToString(w)
}
