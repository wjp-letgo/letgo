package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ImageEntity
type ImageEntity struct {
	ImageUrlList []string `json:"image_url_list"`
	ImageIdList  []string `json:"image_id_list"`
}

//String
func (i ImageEntity) String() string {
	return lib.ObjectToString(i)
}

//TierImageEntity
type TierImageEntity struct {
	ImageID  string `json:"image_id"`
	ImageURL string `json:"image_url"`
}

//String
func (i TierImageEntity) String() string {
	return lib.ObjectToString(i)
}
