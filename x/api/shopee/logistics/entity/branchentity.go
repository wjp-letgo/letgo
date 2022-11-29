package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//BranchEntity
type BranchEntity struct {
	BranchID int64  `json:"branch_id"`
	Region   string `json:"region"`
	State    string `json:"state"`
	City     string `json:"city"`
	Address  string `json:"address"`
	ZipCode  string `json:"zipcode"`
	District string `json:"district"`
	Town     string `json:"town"`
}

//String
func (b BranchEntity) String() string {
	return lib.ObjectToString(b)
}
