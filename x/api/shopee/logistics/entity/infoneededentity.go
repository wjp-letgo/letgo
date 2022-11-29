package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//InfoNeededEntity
type InfoNeededEntity struct {
	Dropoff       []string `json:"dropoff"`
	Pickup        []string `json:"pickup"`
	NonIntegrated []string `json:"non_integrated"`
}

//String
func (i InfoNeededEntity) String() string {
	return lib.ObjectToString(i)
}
