package entity

import (
	"github.com/wjp-letgo/letgo/lib"
	"github.com/wjp-letgo/letgo/x/api/shopee/commonentity"
)

//GetItemBaseInfoResult
type GetItemBaseInfoResult struct {
	commonentity.Result
	Response GetItemBaseInfoResultResponse `json:"response"`
	Warning  string                        `json:"warning"`
}

//String
func (g GetItemBaseInfoResult) String() string {
	return lib.ObjectToString(g)
}

//GetItemBaseInfoResultResponse
type GetItemBaseInfoResultResponse struct {
	ItemList []ItemEntity `json:"item_list"`
}

//String
func (g GetItemBaseInfoResultResponse) String() string {
	return lib.ObjectToString(g)
}
