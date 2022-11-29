package session

import "github.com/wjp-letgo/letgo/lib"

//SessionConfig
type SessionConfig struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Expire int    `json:"expire"`
	Path   string `json:"path"`
	Prefix string `json:"prefix"`
}

//String
func (s SessionConfig) String() string {
	return lib.ObjectToString(s)
}
