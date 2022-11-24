package entity

import (
	"github.com/wjpxxx/letgo/lib"
)

//ImageEntity
type ImageEntity struct{
	ImageIdList []string `json:"image_id_list"`
	ImageUrlList []string `json:"image_url_list"`
}

//String
func(p ImageEntity)String()string{
	return lib.ObjectToString(p)
}

//TierImageEntity
type TierImageEntity struct{
	ImageID string `json:"image_id"`
	ImageURL string `json:"image_url"`
}

//String
func(i TierImageEntity)String()string{
	return lib.ObjectToString(i)
}