package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//VideoInfoEntity
type VideoInfoEntity struct {
	VideoUrl     string `json:"video_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Duration     int    `json:"duration"`
}

//String
func (v VideoInfoEntity) String() string {
	return lib.ObjectToString(v)
}
