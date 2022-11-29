package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//ThumbnailUrlEntity
type ThumbnailUrlEntity struct {
	ImageUrlRegion string `json:"image_url_region"`
	ImageUrl       string `json:"image_url"`
}

//String
func (g ThumbnailUrlEntity) String() string {
	return lib.ObjectToString(g)
}
