package entity

import(
	"github.com/wjpxxx/letgo/x/api/shopee/commonentity"
	"github.com/wjpxxx/letgo/lib"
)

//SetNoteResult
type SetNoteResult struct{
	commonentity.Result
}

//String
func(s SetNoteResult)String()string{
	return lib.ObjectToString(s)
}