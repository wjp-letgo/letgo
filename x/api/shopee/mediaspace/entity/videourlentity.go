package entity

import (
	"github.com/wjp-letgo/letgo/lib"
)

//VideoUrlEntity
type VideoUrlEntity struct{
	VideoUrlRegion string `json:"video_url_region"`
	VideoUrl string `json:"video_url"`
}

//String
func(g VideoUrlEntity)String()string{
	return lib.ObjectToString(g)
}