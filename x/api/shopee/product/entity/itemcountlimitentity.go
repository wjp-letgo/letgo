package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ItemCountLimitEntity
type ItemCountLimitEntity struct {
	MaxLimit int `json:"max_limit"`
}

//String
func (p ItemCountLimitEntity) String() string {
	return lib.ObjectToString(p)
}
