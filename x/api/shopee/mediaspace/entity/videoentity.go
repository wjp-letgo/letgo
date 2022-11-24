package entity

import (
	"github.com/wjpxxx/letgo/lib"
)
//VideoInfoEntity
type VideoInfoEntity struct{
	VideoUrlList []VideoUrlEntity `json:"video_url_list"`
	ThumbnailUrlList []ThumbnailUrlEntity `json:"thumbnail_url_list"`
	Duration int `json:"duration"`
}
//String
func(g VideoInfoEntity)String()string{
	return lib.ObjectToString(g)
}