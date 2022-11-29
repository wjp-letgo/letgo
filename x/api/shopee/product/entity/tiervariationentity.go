package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//TierVariationEntity
type TierVariationEntity struct {
	OptionList []OptionEntity `json:"option_list"`
	Name       string         `json:"name"`
}

//String
func (s TierVariationEntity) String() string {
	return lib.ObjectToString(s)
}
