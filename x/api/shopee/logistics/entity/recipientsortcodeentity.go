package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//RecipientSortCodeEntity
type RecipientSortCodeEntity struct{
	FirstRecipientSortCode string `json:"first_recipient_sort_code"`
	SecondRecipientSortCode string `json:"second_recipient_sort_code"`
	ThirdRecipientSortCode string `json:"third_recipient_sort_code"`
}

//String
func(r RecipientSortCodeEntity)String()string{
	return lib.ObjectToString(r)
}
