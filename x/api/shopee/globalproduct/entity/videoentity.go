package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//VideoEntity
type VideoEntity struct {
	VideoUrl     string `json:"video_url"`
	ThumbnailUrl string `json:"thumbnail_url"`
	Duration     int    `json:"duration"`
}

//String
func (p VideoEntity) String() string {
	return lib.ObjectToString(p)
}
